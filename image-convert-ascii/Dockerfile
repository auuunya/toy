FROM python:3.6.15-slim-buster as running

# Set the working directory in the container
WORKDIR /app

COPY . /app

# Install any needed packages specified in requirements.txt
RUN python -m pip install --upgrade pip && pip install --no-cache-dir -r requirements.txt
CMD ["python", "image_covert_ascii.py", "-i", "bird.jpg", "-t", "dark"]