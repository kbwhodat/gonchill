from selenium_profiles.webdriver import Chrome
from selenium_profiles.profiles import profiles
from seleniumwire import webdriver
import json
import time

def gen_driver():
    try:
        profile = profiles.Windows()
        chrome_options = webdriver.ChromeOptions()
        user_agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"

        chrome_options.add_argument("--start-maximized")
        chrome_options.add_argument("--window-size=1920,1080")
        chrome_options.add_argument(f"--user-agent={user_agent}")
        chrome_options.add_argument("--headless=new")
        chrome_options.add_argument("--no-sandbox")
        chrome_options.add_argument("--disable-dev-shm-usage")
        chrome_options.add_experimental_option("prefs", {
            "profile.default_content_setting_values.popups": 1,  # Allow pop-ups
            "profile.default_content_setting_values.cookies": 1,
            "profile.cookie_controls_mode": 0,
            "profile.block_third_party_cookies": False,
        })


        driver = Chrome(profile, options=chrome_options, uc_driver=False)
        return driver
    except Exception as e:
        print("Error in Driver: ",e)

def main():
    driver = gen_driver()  # Initialize the driver using your function
    try:
        time.sleep(2)
        location = "/tmp/cookies.json"

        driver.get('https://youtube.com')
        time.sleep(2)
        driver.execute_script("window.open('', '_blank');")
        time.sleep(10)

        driver.switch_to.window(driver.window_handles[-1])
        time.sleep(1)

        # user_agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"
        driver.execute_cdp_cmd('Network.setUserAgentOverride', {
        "userAgent": user_agent,
        "platform": "Linux",
        "acceptLanguage": "en-US,en;q=0.9",
        "userAgentMetadata": {
            "brands": [{"brand": "Chromium", "version": "137"}, {"brand": "Not-A.Brand", "version": "99"}],
            "fullVersion": "137.0.0.0",
            "platform": "Linux",
            "platformVersion": "5.10",
            "architecture": "x86",
            "model": "",
            "mobile": False
        }
    })

        driver.execute_script("window.location.href = 'https://en.rarbg-official.com/episodes/true-detective-2014-season-4-episode-1';")


        time.sleep(8)
        driver.switch_to.window(driver.window_handles[1])

        time.sleep(2)


        print("trying to get cf_clearance cookie...")
        cookies = driver.get_cookie("cf_clearance")
        cookie_list = [cookies] if cookies else []
        print("generating cookies...")
        with open(location, 'w') as filehandler:
            json.dump(cookie_list, filehandler, indent=4)


    except Exception as e:
        print("Error in Driver: ",e)

if __name__ == "__main__":
    main()
