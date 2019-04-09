FROM python:3.6-alpine
ADD . /data
WORKDIR /data
RUN pip install --no-cache-dir -r requirements.txt -i https://pypi.doubanio.com/simple
ENTRYPOINT ["python", "main.py"]