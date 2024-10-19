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

# can use .eq after select
# Placeholder code in table and select
#data = supabase.table("todos").select("id, name").execute()
#print(data)

#insert data
#created_at = datetime.utcnow() - timedelta(hours=2)
#data = supabase.table("todos").insert({"name":"Todo 2", "created_at": str(created_at)}).execute()

#update
#data = supabase.table("todos").update({"name": "updated name"}).eq("id", 1).execute()
#data = supabase.table("todos").select("*").execute()
#print(data)

#delete
#data = supabase.table("todos").delete().eq("id", 1).execute()


app = Flask(__name__)

@app.route('/')
def hello():
    return 'Hello World!'
# TODO: Login
# TODO: Log_Out
# TODO: Add_Faculty
# TODO: Add_Patient
# TODO: View_Patients
# TODO: View_Faculty
# TODO: Delete_User
# TODO: Forms?
# TODO: Feedback?
# TODO: IDK

if __name__ == '__main__':
    app.run(debug=True)