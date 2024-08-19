#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   image_convert_ascii.py
@Time    :   2024/07/05 15:51:07
@Desc    :   None
'''

# here put the import lib
from PIL import Image
import argparse


ASCII_CHARS = "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
MAX_PIXEL_SIZE = 255

def read_img_to_martix(path, size=(128,128)):
    if path is None:
        raise Exception("Not Found To Read Image.")
    img = Image.open(path, "r").convert("RGB")
    img.thumbnail(size)
    martix = []
    for y in range(img.height):
        row = []
        for x in range(img.width):
            pixel = img.getpixel((x, y))
            row.append(pixel)
        martix.append(row)
    return martix

def pixel_to_bright(martix, algo_name):
    newline = []
    for lines in martix:
        row = []
        for line in lines:
            if algo_name == 'average':
                intensity = (line[0] + line[1] + line[2]) / 3.0
            elif algo_name == 'max_min':
                intensity = (max(line) + min(line)) / 2.0
            elif algo_name == 'luminosity':
                intensity = 0.21 * line[0] + 0.72 * line[1] + 0.07 * line[2]
            else:
                raise Exception("Unrecognized algo_name: %s" % algo_name)
            row.append(intensity)
        newline.append(row)
    return newline


def covert_to_ascii(bright, ascii_chars, theme="dark"):
    ascii_martix = []
    for row in bright:
        ascii_row = []
        for p in row:
            if theme == "dark":
                ascii_row.append(ascii_chars[int(p / MAX_PIXEL_SIZE * (len(ascii_chars) - 1))])
            elif theme == "light":
                ascii_row.append(ascii_chars[~(int(p / MAX_PIXEL_SIZE * (len(ascii_chars) - 1)))])
            else:
                raise Exception("Unrecognized theme: %s" % theme)
        ascii_martix.append(ascii_row)
    return ascii_martix

def print_to_ascii(ascii_list, bold=1):
    ascii_text = ""
    for line in ascii_list:
        if bold:
            line = [p*bold for p in line]
        ascii_text += "".join(line) + "\n"
    print(ascii_text)

def print_arguments():
    parse = argparse.ArgumentParser(
        description="Image Convert To ASCII Art"
    )
    parse.add_argument(
        "-i", 
        "--image",
        type=str, 
        required=True, 
        help="Path to the input image."
    )
    parse.add_argument(
        "-a", 
        "--alog_name",
        type=str,  
        choices=["average", "max_min", "luminosity"],
        default="average",
        help="Alog Name of the ASCII art."
    )
    parse.add_argument(
        "-t", 
        "--theme",
        type=str, 
        choices=["dark", "light"],
        default="dark",
        help="Theme of the ASCII art (dark or light)."
    )
    parse.add_argument(
        "-s",
        "--size",
        type=int,
        nargs=2,
        metavar=("width", "height"),
        default=(128, 128),
        help="Thumbnail size for the image, specified as width and height."
    )
    parse.add_argument(
        "-b",
        "--bold",
        type=int,
        default=1,
        help="Number of placeholder"
    )
    args = parse.parse_args()
    return args

if __name__ == '__main__':
    args = print_arguments()
    path = args.image
    theme = args.theme
    size = args.size
    bold = args.bold
    alog_name = args.alog_name
    martix = read_img_to_martix(path, size=size)
    newarr = pixel_to_bright(martix, algo_name=alog_name)
    asciirows = covert_to_ascii(newarr, list(ASCII_CHARS), theme=theme)
    print_to_ascii(asciirows, bold)