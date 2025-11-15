import requests
import json
from bs4 import BeautifulSoup

def check_website(url):
    """Checks the status of a website and returns information about it."""

    try:
        response = requests.get(url, timeout=10)  # Set a timeout to prevent indefinite hanging
        response.raise_for_status()  # Raise an exception for bad status codes (4xx or 5xx)

        status_code = response.status_code
        content = response.text

        if status_code == 200:
            soup = BeautifulSoup(content, 'html.parser')
            title = soup.title.string if soup.title else "No Title Found"  # Extract the title

            # Basic keyword check (customize as needed)
            keywords = ["Welcome", "Home", "Product"]  # Example keywords
            page_type = "General Page"  # Default
            for keyword in keywords:
                if keyword.lower() in content.lower():
                    page_type = "Likely a " + keyword + " Page"
                    break

            return {
                "url": url,
                "status": "OK",
                "status_code": status_code,
                "title": title,
                "page_type": page_type  # Added page type
            }

        else:
            return {
                "url": url,
                "status": "Error",
                "status_code": status_code,
                "message": f"HTTP Error: {status_code}"
            }


    except requests.exceptions.RequestException as e:
        return {
            "url": url,
            "status": "Error",
            "message": str(e)  # More descriptive error message
        }
    except Exception as e: # Catching general exceptions
        return {
            "url": url,
            "status": "Error",
            "message": f"An error occurred: {str(e)}"
        }



def check_websites(urls):
    """Checks multiple websites and returns the results as a JSON string."""

    results =
    for url in urls:
        if not url.startswith("http"): #Adding protocol if not present
            url = "https://" + url
        results.append(check_website(url))
    return json.dumps(results, indent=4) # indent for pretty printing


if __name__ == "__main__":
    website_urls = [
        "self-migration-test.pl.s2-eu.capgemini.com"
        "service42.pl.s2-eu.capgemini.com"
        "service42v1.pl.s2-eu.capgemini.com"
        "servicebroker.pl.s2-eu.capgemini.com"
        "sesame.pl.s2-eu.capgemini.com"
        "sesp.pl.s2-eu.capgemini.com"
        "shared-lib-test.pl.s2-eu.capgemini.com"
        "shared-services.pl.s2-eu.capgemini.com"
        "sherlockpl.pl.s2-eu.capgemini.com"
        "siamoi.pl.s2-eu.capgemini.com"
        "si-apa.pl.s2-eu.capgemini.com"
        "sibr.pl.s2-eu.capgemini.com"
        "si-log.pl.s2-eu.capgemini.com"
        "sionisr.pl.s2-eu.capgemini.com"
        "siorg-pl.pl.s2-eu.capgemini.com"
        "sips.pl.s2-eu.capgemini.com"
        "sister.pl.s2-eu.capgemini.com"
        "sky-ai4qa.pl.s2-eu.capgemini.com"
        "skyde.pl.s2-eu.capgemini.com"
        "smcp.pl.s2-eu.capgemini.com"
        "sncfgda.pl.s2-eu.capgemini.com"
        "sncf-mat-b2b.pl.s2-eu.capgemini.com"
        "sncfrx-pss-iot.pl.s2-eu.capgemini.com"
        "soc-toulouse.pl.s2-eu.capgemini.com"
        "spf4dusi.pl.s2-eu.capgemini.com"
        "sqgenaiedf.pl.s2-eu.capgemini.com"
        "ssb-rein-ima.pl.s2-eu.capgemini.com"
        "stimecdship.pl.s2-eu.capgemini.com"
        "stsi.pl.s2-eu.capgemini.com"
        "swc.s2-eu.capgemini.com"
        "sw-hub.pl.s2-eu.capgemini.com"
        "switch.pl.s2-eu.capgemini.com"
        "systeamuo4-6.pl.s2-eu.capgemini.com"
        "tarros.pl.s2-eu.capgemini.com"
        "tas.pl.s2-eu.capgemini.com"
        "tech-it-up.pl.s2-eu.capgemini.com"
        "temp-pl.pl.s2-eu.capgemini.com"
        "testansible.pl.s2-eu.capgemini.com"
        "test-merge-pl.pl.s2-eu.capgemini.com"
        "test-new-nfs-n.pl.s2-eu.capgemini.com"
        "test-new-nfs.pl.s2-eu.capgemini.com"
        "test-newnfs.pl.s2-eu.capgemini.com"
        "tewatnext.pl.s2-eu.capgemini.com"
        "tgits-4biz.pl.s2-eu.capgemini.com"
        "tim-whs.pl.s2-eu.capgemini.com"
        "tma-anses-pl.pl.s2-eu.capgemini.com"
        "tma-dracar.pl.s2-eu.capgemini.com"
        "tmainwiwina.pl.s2-eu.capgemini.com"
        "tmaoss.s2-eu.capgemini.com"
        "tma-safran.pl.s2-eu.capgemini.com"
        "tmasafran.pl.s2-eu.capgemini.com"
        "tma-sitere.pl.s2-eu.capgemini.com"
        "tmatotal.pl.s2-eu.capgemini.com"
        "tma-uimm.pl.s2-eu.capgemini.com"
        "tnard-dlrbdx.pl.s2-eu.capgemini.com"
        "tno.pl.s2-eu.capgemini.com"
        "toulousepl.pl.s2-eu.capgemini.com"
        "tpoffice.pl.s2-eu.capgemini.com"
        "tsnindus.pl.s2-eu.capgemini.com"
        "udti.pl.s2-eu.capgemini.com"
        "unode50.pl.s2-eu.capgemini.com"
        "urssaf-raf.pl.s2-eu.capgemini.com"
        "uwvupa.s2-eu.capgemini.com"
        "vbto.pl.s2-eu.capgemini.com"
        "vcx.pl.s2-eu.capgemini.com"
        "vegeo-dev.pl.s2-eu.capgemini.com"
        "ventasprp.pl.s2-eu.capgemini.com"
        "vfeompl.pl.s2-eu.capgemini.com"
        "vf-gea-nva.pl.s2-eu.capgemini.com"
        "vf-gl-build.pl.s2-eu.capgemini.com"
        "visitecdsedfar.pl.s2-eu.capgemini.com"
        "wallbox.pl.s2-eu.capgemini.com"
        "whirlpool-it.pl.s2-eu.capgemini.com"
        "wind3cloud.pl.s2-eu.capgemini.com"
        "wind3-cma.pl.s2-eu.capgemini.com"
        "wind3-frodi.pl.s2-eu.capgemini.com"
        "wind-mag.pl.s2-eu.capgemini.com"
        "wind-ram.pl.s2-eu.capgemini.com"
        "worem.pl.s2-eu.capgemini.com"
        "xtech.pl.s2-eu.capgemini.com"
        "zephyr-heol.pl.s2-eu.capgemini.com"
    ]

    json_output = check_websites(website_urls)
    print(json_output)