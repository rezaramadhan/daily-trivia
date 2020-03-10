# Use an official Python runtime as a parent image
FROM python:3.7
LABEL maintainer="rezaramadhan.m@gmail.com"

# Set environment varibles
ENV PYTHONUNBUFFERED 1
ENV FLASK_APP daily_trivia.py
ENV FLASK_ENV development

COPY ./requirements.txt /app/requirements.txt

# Install any needed packages specified in requirements.txt
RUN  pip install --upgrade pip && pip install -r /app/requirements.txt

# Copy the current directory contents into the container at /code/
COPY ./templates/ /app/templates
COPY ./daily_trivia.py /app/daily_trivia.py

# Set the working directory to /code/
WORKDIR /app

EXPOSE 5000

CMD flask run --host 0.0.0.0 --port 5000
