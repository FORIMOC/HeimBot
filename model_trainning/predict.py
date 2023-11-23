import torch
import torch.nn as nn
import pandas as pd
from transformers import BertTokenizer, BertModel
from torch.utils.data import DataLoader

# 本地预训练模型路径
BERT_PATH = 'bert-base-chinese'

device = 'cuda' if torch.cuda.is_available() else 'cpu'

batch_size = 5

# 分词器
tokenizer = BertTokenizer.from_pretrained(BERT_PATH)


# BERT 模型
class BertClassifier(nn.Module):
    def __init__(self, dropout=0.5):
        super(BertClassifier, self).__init__()
        self.bert = BertModel.from_pretrained(BERT_PATH)
        self.dropout = nn.Dropout(dropout)
        self.linear = nn.Linear(768, 2)
        self.relu = nn.ReLU()

    def forward(self, input_id, mask):
        _, pooled_output = self.bert(input_ids=input_id, attention_mask=mask, return_dict=False)
        dropout_output = self.dropout(pooled_output)
        linear_output = self.linear(dropout_output)
        final_layer = self.relu(linear_output)
        return final_layer


model = BertClassifier()
model.load_state_dict(torch.load('model/bert_9.pt'))
model = model.to(device)
model.eval()


def preprocess_text(text):
    encoded_input = tokenizer.encode_plus(
        text,
        add_special_tokens=True,
        padding='max_length',
        max_length=128,
        return_tensors='pt'
    )
    input_id = encoded_input['input_ids']
    mask = encoded_input['attention_mask']
    return input_id, mask


def predict_batch(texts):
    input_ids = []
    masks = []
    for text in texts:
        input_id, mask = preprocess_text(text)
        input_ids.append(input_id)
        masks.append(mask)
    input_ids = torch.cat(input_ids, dim=0)
    masks = torch.cat(masks, dim=0)
    dataset = torch.utils.data.TensorDataset(input_ids, masks)
    dataloader = DataLoader(dataset, batch_size=batch_size)
    predictions = []
    for batch in dataloader:
        batch = tuple(t.to(device) for t in batch)
        input_id, mask = batch
        with torch.no_grad():
            outputs = model(input_id.to(device), mask.to(device))
            _, predicted_labels = torch.max(outputs, dim=1)
            predictions.extend(predicted_labels.cpu().numpy())
    return predictions


with open('data/test.csv', 'r', encoding='utf8') as f:
    texts = f.readlines()
predictions = predict_batch(texts)
print(predictions)
