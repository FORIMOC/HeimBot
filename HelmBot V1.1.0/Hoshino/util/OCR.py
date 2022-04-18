import sys
import pytesseract
import requests
from PIL import Image
from io import BytesIO
# pytesseract.pytesseract.tesseract_cmd = r'C:\Program Files\Tesseract-OCR\tesseract.exe'
# tessdata_dir_config = r'--tessdata-dir "C:\Program Files\Tesseract-OCR\tessdata"'
pytesseract.pytesseract.tesseract_cmd = r'/usr/local/bin/tesseract'
tessdata_dir_config = r'--tessdata-dir "/usr/local/share/tessdata"'


img_src = sys.argv[1]
resp = requests.get(img_src)
img = Image.open(BytesIO(resp.content))
print(pytesseract.image_to_string(img, lang=sys.argv[2], config=tessdata_dir_config))
