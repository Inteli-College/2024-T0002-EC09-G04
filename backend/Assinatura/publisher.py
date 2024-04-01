import time
import random
import paho.mqtt.client as paho
from paho import mqtt
from dotenv import load_dotenv
import os
import json
from Crypto.PublicKey import RSA
from Crypto.Signature import pkcs1_15
from Crypto.Hash import SHA256

load_dotenv()

# MQTT Broker settings
broker_address = os.getenv("BROKER_ADDRESS")
port = 8883
topic = "data/sensor1"
username = os.getenv("USER_NAME")
password = os.getenv("PASSWORD")

# Load RSA private key
with open('private_key.pem', 'r') as f:
    private_key = f.read()

# Function to simulate sensor readings
def generate_sensor_data():
    co_reading = random.uniform(1, 1000)
    no2_reading = random.uniform(0.05, 10)
    ethanol_reading = random.uniform(10, 500)
    hydrogen_reading = random.uniform(1, 1000)
    ammonia_reading = random.uniform(1, 500)

    return {
        "CO": co_reading,
        "NO2": no2_reading,
        "Ethanol": ethanol_reading,
        "Hydrogen": hydrogen_reading,
        "Ammonia": ammonia_reading
    }

# Assinar dados
def sign_data(data, private_key):
    key = RSA.import_key(private_key)
    h = SHA256.new(data.encode())
    signature = pkcs1_15.new(key).sign(h)
    return signature

# MQTT setup and connection
client = paho.Client(paho.CallbackAPIVersion.VERSION2, "Publisher",
                     protocol=paho.MQTTv5)

def on_connect(client, userdata, flags, reason_code, properties):
    print(f"CONNACK received with code {reason_code}")
    client.subscribe(topic, qos=1)

client.on_connect = on_connect
client.tls_set(tls_version=mqtt.client.ssl.PROTOCOL_TLS)
client.username_pw_set(username, password)
client.connect(broker_address, port, 60)

try:
    # Start the simulation
    client.loop_start()
    while True:
        sensor_data = generate_sensor_data()
        sensor_data_json = json.dumps(sensor_data)
        signature = sign_data(sensor_data_json, private_key)
        message = {
            "data": sensor_data_json,
            "signature": signature.hex()
        }
        message_json = json.dumps(message)
        client.publish(topic, message_json, qos=1, retain=False)
        time.sleep(5)
except KeyboardInterrupt:
    # Gracefully handle interrupt (Ctrl+C) to disconnect from MQTT broker
    client.disconnect()
    print("Simulation stopped.")
