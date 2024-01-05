#encoding='utf-8'
"""
@author=wanggang
"""
"""
import requests

def retrieve_all_transactions(punk_address, api_key, start_block=0, end_block=99999999, page=1):
    etherscan_url = f'http://api.etherscan.io/api?module=account&action=txlist&address={punk_address}&startblock={start_block}&endblock={end_block}&sort=asc&page={page}&apikey={api_key}'
    count = 0  # 计数器
    try:
        response = requests.get(etherscan_url)
        if response.status_code == 200:
            data = response.json()
            if data['status'] == "1":
                transactions = data['result']
                if len(transactions) > 0:
                    for transaction in transactions:
                        count+=1
                        print(count,'--------',transaction)
                    # 继续获取下一页数据
                    page += 1
                    retrieve_all_transactions(punk_address, api_key, start_block, end_block, page)
                else:
                    print("No more transactions.")
            else:
                print(f'Etherscan API error: {data["message"]}')
        else:
            print(f'HTTP request failed with status code: {response.status_code}')
    except requests.exceptions.RequestException as e:
        print(f'HTTP request error: {e}')

# 设置要检索的地址和API密钥
punk_address = '0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb'
api_key = 'IUXQNAEUS1WYN9ITW4J23926R4SIDZ5C5J'

# 调用函数开始检索
retrieve_all_transactions(punk_address, api_key)
"""
import requests

# def retrieve_all_transactions(punk_address, api_key, start_block=0, end_block=99999999, page=1):
#     etherscan_url = f'http://api.etherscan.io/api?module=account&action=txlist&address={punk_address}&startblock={start_block}&endblock={end_block}&sort=asc&page={page}&apikey={api_key}'
#     total_transactions = 0
#     all_transactions = []  # 用于保存所有事务的列表
#
#     try:
#         response = requests.get(etherscan_url)
#         if response.status_code == 200:
#             data = response.json()
#             if data['status'] == "1":
#                 transactions = data['result']
#                 if len(transactions) > 0:
#                     for transaction in transactions:
#                         total_transactions += 1
#                         all_transactions.append(transaction)  # 将事务添加到列表中
#                     # 继续获取下一页数据
#                     page += 1
#                     print(f'Total transactions so far: {total_transactions}')
#                     retrieve_all_transactions(punk_address, api_key, start_block, end_block, page)
#                 else:
#                     print("No more transactions.")
#                     print(f'Total transactions: {total_transactions}')
#                     # 处理所有事务
#                     for i, transaction in enumerate(all_transactions):
#                         print(f'Transaction {i + 1}: {transaction}')
#             else:
#                 print(f'Etherscan API error: {data["message"]}')
#         else:
#             print(f'HTTP request failed with status code: {response.status_code}')
#     except requests.exceptions.RequestException as e:
#         print(f'HTTP request error: {e}')
#
# # 设置要检索的地址和API密钥
# #punk_address = 'YOUR_PUNK_ADDRESS'
# #api_key = 'YOUR_API_KEY'
# punk_address = '0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb'
# api_key = 'IUXQNAEUS1WYN9ITW4J23926R4SIDZ5C5J'
# # 调用函数开始检索
# retrieve_all_transactions(punk_address, api_key)
"""
import requests

def retrieve_all_transactions(punk_address, api_key, start_block=0, end_block=99999999, page=1, total_transactions=0):
    etherscan_url = f'http://api.etherscan.io/api?module=account&action=txlist&address={punk_address}&startblock={start_block}&endblock={end_block}&sort=asc&page={page}&apikey={api_key}'
    all_transactions = []  # 用于保存所有事务的列表

    try:
        response = requests.get(etherscan_url)
        if response.status_code == 200:
            data = response.json()
            if data['status'] == "1":
                transactions = data['result']
                if len(transactions) > 0:
                    for transaction in transactions:
                        total_transactions += 1
                        all_transactions.append(transaction)  # 将事务添加到列表中
                    # 继续获取下一页数据
                    page += 1
                    print(f'Total transactions so far: {total_transactions}')
                    retrieve_all_transactions(punk_address, api_key, start_block, end_block, page, total_transactions)
                else:
                    print("No more transactions.")
                    print(f'Total transactions: {total_transactions}')
                    # 处理所有事务
                    for i, transaction in enumerate(all_transactions):
                        print(f'Transaction {i + 1}: {transaction}')
            else:
                print(f'Etherscan API error: {data["message"]}')
        else:
            print(f'HTTP request failed with status code: {response.status_code}')
    except requests.exceptions.RequestException as e:
        print(f'HTTP request error: {e}')

# 设置要检索的地址和API密钥
punk_address = '0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb'
api_key = 'IUXQNAEUS1WYN9ITW4J23926R4SIDZ5C5J'

# 调用函数开始检索
retrieve_all_transactions(punk_address, api_key)
"""
# import requests
#
# def retrieve_all_transactions(punk_address, api_key, start_block=0, end_block=99999999):
#     page = 1
#     all_transactions = []
#
#     while True:
#         etherscan_url = f'http://api.etherscan.io/api?module=account&action=txlist&address={punk_address}&startblock={start_block}&endblock={end_block}&sort=asc&page={page}&apikey={api_key}'
#
#         try:
#             response = requests.get(etherscan_url)
#             if response.status_code == 200:
#                 data = response.json()
#                 if data['status'] == "1":
#                     transactions = data['result']
#                     if len(transactions) > 0:
#                         all_transactions.extend(transactions)
#                         print(f'Fetched {len(transactions)} transactions on page {page}')
#                         page += 1
#                     else:
#                         print("No more transactions.")
#                         break
#                 else:
#                     print(f'Etherscan API error: {data["message"]}')
#                     break
#             else:
#                 print(f'HTTP request failed with status code: {response.status_code}')
#                 break
#         except requests.exceptions.RequestException as e:
#             print(f'HTTP request error: {e}')
#             break
#
#     return all_transactions
#
# # 设置要检索的地址和API密钥
# # punk_address = 'YOUR_PUNK_ADDRESS'
# # api_key = 'YOUR_API_KEY'
# punk_address = '0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb'
# api_key = 'IUXQNAEUS1WYN9ITW4J23926R4SIDZ5C5J'
# # 调用函数开始检索
# all_transactions = retrieve_all_transactions(punk_address, api_key)
# print(f'Total transactions: {len(all_transactions)}')
#
# # 现在，all_transactions 中包含了所有的事务数据
import requests
import pandas as pd

def retrieve_transactions(punk_address, api_key, start_block=0, end_block=99999999, max_pages=2):
    page = 1
    all_transactions = []

    while page <= max_pages:
        etherscan_url = f'http://api.etherscan.io/api?module=account&action=txlist&address={punk_address}&startblock={start_block}&endblock={end_block}&sort=asc&page={page}&apikey={api_key}'

        try:
            response = requests.get(etherscan_url)
            if response.status_code == 200:
                data = response.json()
                if data['status'] == "1":
                    transactions = data['result']
                    if len(transactions) > 0:
                        all_transactions.extend(transactions)
                        print(f'Fetched {len(transactions)} transactions on page {page}')
                        page += 1
                    else:
                        print("No more transactions.")
                        break
                else:
                    print(f'Etherscan API error: {data["message"]}')
                    break
            else:
                print(f'HTTP request failed with status code: {response.status_code}')
                break
        except requests.exceptions.RequestException as e:
            print(f'HTTP request error: {e}')
            break

    return all_transactions

# 设置要检索的地址和API密钥
punk_address = '0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb'
api_key = 'IUXQNAEUS1WYN9ITW4J23926R4SIDZ5C5J'
# 调用函数开始检索前两页的事务
all_transactions = retrieve_transactions(punk_address, api_key, max_pages=2)

# # 打印前两页的事务数据
# for i, transaction in enumerate(all_transactions):
#     print(f'Transaction {i + 1}: {transaction}')
# 将事务数据转换为pandas DataFrame
df = pd.DataFrame(all_transactions)

# 保存数据到Excel文件
df.to_excel('transaction_data.xlsx', index=False)