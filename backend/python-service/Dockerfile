FROM python:3.9.18-slim

WORKDIR /app

ADD . /app

RUN pip install --trusted-host pypi.python.org Flask

EXPOSE 5000

CMD ["python", "app.py"]

