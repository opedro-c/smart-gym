import argparse
import time
import random
import json
import paho.mqtt.client as mqtt

def publish_messages(host: str, port: int, topic: str, user_id: int, origin_id: str, start: int) -> None:
    client = mqtt.Client()
    client.connect(host, port, 60)

    THREE_MINUTES = 180000
    end = start + THREE_MINUTES

    current_time = start
    while current_time <= end:
        payload = {
            "user_id": user_id,
            "origin_id": origin_id,
            "data": {
                "started_at": current_time,
                "finished_at": current_time + 60,
                "weight": random.uniform(5, 20),
            },
        }

        client.publish(topic, json.dumps(payload))
        print(f"Published message: {json.dumps(payload)}")

        current_time += 600  # Increment by near 1 minute
        time.sleep(0.1)  # Small delay to avoid flooding the broker

    client.disconnect()

def main() -> None:
    parser = argparse.ArgumentParser(description="Publish MQTT messages with random weight data.")
    parser.add_argument("--host", type=str, default='localhost', help="MQTT broker address")
    parser.add_argument("--port", type=int, default=1883, help="MQTT broker port")
    parser.add_argument("--topic", type=str, default='pedroc_aragao@outlook.com/exercise', help="MQTT topic to publish to")
    parser.add_argument("--user_id", type=int, required=True, help="User ID")
    parser.add_argument("--origin_id", type=str, required=True, help="Origin ID")
    parser.add_argument("--start", type=int, required=True, help="Start timestamp (Unix epoch)")

    args = parser.parse_args()

    publish_messages(
        host=args.host,
        port=args.port,
        topic=args.topic,
        user_id=args.user_id,
        origin_id=args.origin_id,
        start=args.start,
    )

if __name__ == "__main__":
    main()
