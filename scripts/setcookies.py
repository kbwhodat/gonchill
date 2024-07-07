from selenium_profiles.webdriver import Chrome
from selenium_profiles.profiles import profiles
from seleniumwire import webdriver
import os
import json
import time

def gen_driver():
    try:
        profile = profiles.Windows()
        chrome_options = webdriver.ChromeOptions()
        user_agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

        chrome_options.add_argument("--start-maximized")
        chrome_options.add_argument("--window-size=1920,1080")
        chrome_options.add_argument(f"--user-agent={user_agent}")
        chrome_options.add_argument("--headless=new")
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
        location = "scripts/cookies.json"

        driver.get('https://youtube.com')
        time.sleep(2)
        driver.execute_script("window.open('https://en.rarbg-official.com/episodes/true-detective-2014-season-4-episode-1','_blank');")


        time.sleep(2)
        driver.switch_to.window(driver.window_handles[1])

        time.sleep(5)

        cookies = driver.get_cookie("cf_clearance")
        cookie_list = [cookies] if cookies else []
        with open(location, 'w') as filehandler:
            json.dump(cookie_list, filehandler, indent=4)

            
    except Exception as e:
        print("Error in Driver: ",e)

if __name__ == "__main__":
    main()

