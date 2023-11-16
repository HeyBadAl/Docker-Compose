FROM python:3.9.18-slim

WORKDIR /app

ADD . /app

RUN pip install --trusted-host pypi.python.org Flask pytest

EXPOSE 5000

CMD ["python", "app.py"]

