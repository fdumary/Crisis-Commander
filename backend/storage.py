from dotenv import load_dotenv
load_dotenv()
import os
from flask import Flask
from datetime import datetime, timedelta
from supabase import create_client

# Database config
url = os.environ.get("SUPABASE_URL")
key = os.environ.get("SUPABASE_KEY")

supabase = create_client(url, key)

#resp = supabase.storage().from_("image-bucket").get_public_url("pat_twitter.png")
#print(resp)
resp = supabase.storage().from_("image-bucket").upload("banana.jpeg", "banana.jpeg",
                                                       {"content-type": "image/jpeg"})
print(resp)