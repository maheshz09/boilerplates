
echo "website list written on website.txt file"

if [[ ! -f /tmp/website.txt ]]; then
    echo "Error: File /tmp/website.txt not found!"
    exit 1
fi

while IFS= read -r website; do

    echo "----------------------"
  echo "Checking website: $website"
    echo "----------------------"

    # Fetch and display HTTP headers and specific fields
    curl -ILs "$website" | head -6 #| egrep "HTTP|Content-Type|Content-Security-Policy|Server|Cookie"
    content_type=$(echo "$headers" | grep -i "Content-Type" || echo "Content-Type: Not Found")
    content_security=$(echo "$headers" | grep -i "Content-Security-Policy" || echo "Content-Security-Policy: Not Found")

    echo $headers
    echo "----------------------"
    sleep 1
done < /tmp/website.txt
