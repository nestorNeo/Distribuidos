#/usr/bin/python3

import requests

def main():
    resp = requests.get("http://192.168.56.1:5000/click")
    resp.raise_for_status()

    print(resp.json())

if __name__ == '__main__':
    main()
