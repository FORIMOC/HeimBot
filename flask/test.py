import re
import json

def convert_to_json(text):
    pattern = r'(?P<sender>user\d+)\((?P<user_id>\d+)\):\s*(?P<message>.*?)\s*$'
    messages = re.findall(pattern, text, re.MULTILINE)

    data = []
    group_id = 104  # You can set your desired group_id here

    for sender, user_id, message in messages:
        data.append({
            "group_id": group_id,
            "message": message,
            "post_type": "message",
            "sender": {
                "nickname": sender
            },
            "user_id": int(user_id)
        })

    return {
        "data": data
    }

def main(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as file:
        text = file.read()

    json_data = convert_to_json(text)

    with open(output_file, 'w', encoding='utf-8') as outfile:
        json.dump(json_data, outfile, indent=2, ensure_ascii=False)

if __name__ == "__main__":
    input_file = "input.txt"  # Replace with the path to your input text file
    output_file = "output.json"  # Replace with the desired output JSON file path
    main(input_file, output_file)
