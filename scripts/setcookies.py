from selenium_driverless import webdriver
from selenium_driverless.types.by import By
import json
import asyncio

# note to author: if I forget why this fails make sure to double check the series.go and movies.go files userAgent variable

async def main():

    chrome_options = webdriver.ChromeOptions()
    chrome_options.add_argument("--start-maximized")
    chrome_options.add_argument("--window-size=1920,1080")
    chrome_options.add_argument("--headless=new")
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_options.add_argument("--disable-features=DisableLoadExtensionCommandLineSwitch")
    chrome_options.add_experimental_option("prefs", {
        "profile.default_content_setting_values.popups": 0,  # Allow pop-ups
        "profile.default_content_setting_values.cookies": 1,
        "profile.cookie_controls_mode": 0,
        "profile.block_third_party_cookies": False,
    })

    async with webdriver.Chrome(options=chrome_options) as driver:

        await driver.get('https://en.rarbg-official.is/series', wait_load=True)
        await driver.sleep(0.5)
        await driver.wait_for_cdp("Page.domContentEventFired", timeout=15)
        await driver.sleep(5)
        
        
        location = "/tmp/cookies.json"

        cookies = await driver.get_cookie("cf_clearance")
        print(cookies)

        cookie_list = [cookies] if cookies else []
        print("generating cookies...")
        with open(location, 'w') as filehandler:
            json.dump(cookie_list, filehandler, indent=4)


asyncio.run(main())
