import paho.mqtt.client as paho
from paho import mqtt
from dotenv import load_dotenv
import os
import sqlite3
import datetime
import json
from Crypto.PublicKey import RSA
from Crypto.Signature import pkcs1_15
from Crypto.Hash import SHA256
import binascii

conn = sqlite3.connect('data.db')
c = conn.cursor()
c.execute('''CREATE TABLE IF NOT EXISTS sensor_data (timestamp DATETIME, CO REAL, NO2 REAL, Ethanol REAL, Hydrogen REAL, Ammonia REAL)''')
conn.commit()

load_dotenv()

# MQTT Broker settings
broker_address = os.getenv("BROKER_ADDRESS")
port = 8883
topic = "data/sensor1"
username = os.getenv("USER_NAME")
password = os.getenv("PASSWORD")

# Load RSA public key
with open('public_key.pem', 'r') as f:
    public_key = f.read()

# Callback when a message is received
def on_message(client, userdata, msg):
    msg_data = json.loads(msg.payload.decode())
    data = msg_data["data"]
    signature_hex = msg_data["signature"]
    
    # Verificar assinatura
    def verify_signature(data, signature_hex, public_key):
        signature = bytes.fromhex(signature_hex)
        key = RSA.import_key(public_key)
        h = SHA256.new(data.encode())
        try:
            pkcs1_15.new(key).verify(h, signature)
            return True
        except (ValueError, TypeError, pkcs1_15.VerificationError):
            return False

    if verify_signature(data, signature_hex, public_key):
        sensor_data = json.loads(data)
        CO = sensor_data['CO']
        NO2 = sensor_data['NO2']
        Ethanol = sensor_data['Ethanol']
        Hydrogen = sensor_data['Hydrogen']
        Ammonia = sensor_data['Ammonia']
        
        # get the current time  
        dateTime = datetime.datetime.now()
        
        # Use placeholders in the SQL query
        c.execute("INSERT INTO sensor_data (timestamp, CO, NO2, Ethanol, Hydrogen, Ammonia) VALUES (?, ?, ?, ?, ?, ?)",
                  (dateTime, CO, NO2, Ethanol, Hydrogen, Ammonia))
        conn.commit()
        print("Data inserted into the database.")
    else:
        print("Invalid signature. Data discarded.")

client = paho.Client(paho.CallbackAPIVersion.VERSION2, "Subscriber",
                     protocol=paho.MQTTv5)

def on_connect(client, userdata, flags, reason_code, properties):
    print(f"CONNACK received with code {reason_code}")
    client.subscribe(topic, qos=1)

client.on_connect = on_connect
client.tls_set(tls_version=mqtt.client.ssl.PROTOCOL_TLS)
client.username_pw_set(username, password)
client.on_message = on_message
client.connect(broker_address, port, 60)

try:
    # Subscribe to the topic
    client.subscribe(topic)
    # Start the MQTT loop to receive messages
    client.loop_forever()
except KeyboardInterrupt:
    # Gracefully handle interrupt (Ctrl+C) to disconnect from MQTT broker
    client.disconnect()
    conn.close()
    print("\nSubscriber stopped.")
