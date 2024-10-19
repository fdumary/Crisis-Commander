from dotenv import load_dotenv
from postgrest import APIError

load_dotenv()
import os
from flask import Flask
from datetime import datetime, timedelta
from supabase import create_client

# Database config
# from gotrue.exceptions import APIError
url = os.environ.get("SUPABASE_URL")
key = os.environ.get("SUPABASE_KEY")
supabase = create_client(url, key)

# Go to doc for auth and comment out user after putting right info for user
email: str = "something@supamail.com"
password: str = "something"
user = supabase.auth.sign_up(email=email, password=password)
session = None

data = supabase.table("todos").select("*").execute()
print("Data before sign in: ", data)

"""try:
    session = supabase.auth.sign_in(email=email, password=password)
except APIError:
    print("Login failed")
print(session)
print(session.access_token)"""
# supabase.postgrest.auth(session.access_token)

"""data = supabase.table("todos").select("*").execute()
print("Data before sign in: ", data)"""

# Do policies

# supabase.auth.sign_out()