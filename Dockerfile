FROM python:3.6-alpine
ADD . /data
WORKDIR /data
RUN pip install --no-cache-dir -r requirements.txt
ENTRYPOINT ["python", "main.py"]