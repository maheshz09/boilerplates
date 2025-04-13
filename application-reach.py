import requests

def check_websites(file_path):
    try:
        with open(file_path, 'r') as file:
            websites = file.readlines()
        
        for website in websites:
            website = website.strip()
            if website:  # Ignore empty lines
                try:
                    response = requests.get(website, timeout=10)
                    print(f"Website: {website} | Status Code: {response.status_code} | {'UP' if response.ok else 'DOWN'}")
                except requests.exceptions.RequestException as e:
                    print(f"Website: {website} | Error: {str(e)} | DOWN")
    except FileNotFoundError:
        print(f"File '{file_path}' not found.")
    except Exception as e:
        print(f"An unexpected error occurred: {str(e)}")

# File path to the text file containing website URLs
file_path = 'websites.txt'

# Run the checker
check_websites(file_path)
