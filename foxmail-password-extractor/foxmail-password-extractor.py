#!/bin/env python3

import os, re
import struct
import winreg


def read_registry_key(registryKey, key_path, value_name):
    '''
    读取注册表的值
    '''
    result = ""
    try:
        key = winreg.OpenKey(registryKey, key_path)
        try:
            result, reg_type = winreg.QueryValueEx(key, value_name)
        except FileNotFoundError:
            pass
        winreg.CloseKey(key)
    except FileNotFoundError:
        pass
    return result

def get_foxmail_dir():
    '''
    从注册表中读取foxmail路径
    '''
    value = read_registry_key(winreg.HKEY_CLASSES_ROOT, "Applications\\Foxmail.exe\\shell\\open\\command", "")
    index = value.rfind('\\')
    if index == -1:
        return ""
    value = value[:index + 1]
    return value

def read_file(path):
    with open(path, 'rb') as file:
        return file.read()

def get_client_type(content):
    if content[0] == 0xD0:
        return 'V6'
    elif content[0] == 0x52:
        return 'V7'
    else:
        return 'UNKNOWN'

def find_password(content):
    key = b"Password"
    latest = None
    content_len = len(content)
    key_len = len(key)
    
    for i in range(1, content_len - key_len):
        if content[i - 1] == 0x00 and content[i:i + key_len] == key:
            # Skip "Password" and an unknown int32
            offset = i + key_len + 4
            password_len = struct.unpack('<I', content[offset:offset + 4])[0]
            latest = content[offset + 4:offset + 4 + password_len]
    
    if latest is None:
        raise ValueError("Cannot find password in this file")
    
    return latest

def decrypt_password(client_type: str, encrypted_password_hex: str) -> str:
    """
    解密密码。
    """
    # 定义不同客户端版本的解密密钥
    keys = {
        'V6': b"~draGon~",
        'V7': b"~F@7%m$~"
    }
    
    key = keys.get(client_type)
    if key is None:
        raise ValueError("未知的客户端类型")
    
    # 将十六进制字符串转换为字节数组
    encrypted_bytes = bytearray()
    while encrypted_password_hex:
        byte_value = int(encrypted_password_hex[:2], 16)
        encrypted_bytes.append(byte_value)
        encrypted_password_hex = encrypted_password_hex[2:]
    
    # 解密逻辑
    key_sum = sum(key) % 255
    encrypted_bytes[0] ^= key_sum
    
    decrypted_bytes = bytearray(len(encrypted_bytes) - 1)
    for i in range(len(decrypted_bytes)):
        decrypted_bytes[i] = encrypted_bytes[i + 1] ^ key[i % len(key)]
    
    final_bytes = bytearray(len(decrypted_bytes))
    for i in range(len(final_bytes)):
        if decrypted_bytes[i] < encrypted_bytes[i]:
            final_bytes[i] = 0xFF - encrypted_bytes[i] + decrypted_bytes[i]
        else:
            final_bytes[i] = decrypted_bytes[i] - encrypted_bytes[i]
    
    return final_bytes.decode(errors='ignore')

def sum_bytes(input_bytes):
    return sum(input_bytes)

def get_account_path(account_name):
    """根据账户名获取 Account.rec0 文件路径"""
    foxmail_dir = get_foxmail_dir()
    return os.path.join(foxmail_dir, "Storage", account_name, "Accounts", "Account.rec0")

def extract_emails_from_paths(paths):
    """
    从路径列表中提取每个邮箱地址。
    """
    emails = []
    for path in paths:
        # 使用正则表达式提取邮箱部分
        match = re.search(r'Storage\\([^\\]+)\\', path)
        if match:
            email = match.group(1)
            emails.append(email)
    return emails

def get_email_accounts():
    foxmail_dir = get_foxmail_dir()
    storage_list_file = foxmail_dir + "FMStorage.list"
    if not os.path.exists(storage_list_file):
        raise ValueError("not found storage.list file")
    storages = []
    with open(storage_list_file, 'r', encoding='utf-16') as file:
        for line in file.readlines():
            if not line.strip():
                continue
            storages.append(foxmail_dir + line.strip())
    return storages

def foxmail_password_recovery():
    accounts = get_email_accounts()
    emails = extract_emails_from_paths(accounts)
    passwd_map = {}
    for email in emails:
        account_path = get_account_path(email)

        if not os.path.exists(account_path):
            print(f"{email} Account.rec0 file not found.")
            continue
        content = read_file(account_path)
        client_type = get_client_type(content)
        password_hex = find_password(content)
        password = decrypt_password(client_type, password_hex)
        passwd_map[email] = password
    return passwd_map

if __name__ == "__main__":
    foxmail_user = foxmail_password_recovery()
    print (f"foxmail_user: {foxmail_user}")