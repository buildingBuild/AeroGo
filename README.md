<h2 align="center"> ✈️ AEROGO </h2>

## What it is

An event-driven system built with Apache Kafka and Go that processes streaming flight data and delivers real-time alerts based on state changes.
The system ingests continuous flight updates, detects meaningful transitions (e.g., delay increases, gate changes), and notifies users via SMS with low latency.
(check out the system architecture for more)

## What it does

<li>Allows users to track specific flights (e.g., AA123)</li>
<li>Processes a continuous stream of flight updates</li>
<li>Detects meaningful state changes such as:</li>
<li>Delay increases</li>
<li>Gate changes</li>
<li>Departure events</li>
<li>Sends SMS alerts when changes occur</li>
<li>Prevents duplicate or excessive notifications through controlled processing</li>

## Tech Stack

<li>Language: Go</li>
<li>Streaming: Apache Kafka</li>
<li>Database: PostgreSQL</li>
<li>Messaging: Twilio (SMS)</li>
<li>Architecture: Event-driven</li>

---

## How it works

1. Users subscribe to a flight (`user_id`, `flight_id`)  
2. A polling service continuously fetches flight updates from an external API  
3. Each update is published to Kafka as an event  
4. Kafka stores events in an ordered log partitioned by `flight_id`  
5. A Go-based consumer service reads the event stream  
6. The system compares incoming events against previous state  
7. On detecting meaningful changes:
<li>Queries users tracking the flight</li>
<li>Sends SMS notifications</li>


## Safeguards

<li>Idempotent processing prevents duplicate alerts under Kafka's at-least-once delivery</li>
<li>Retry logic uses exponential backoff for failed SMS deliveries</li>
<li>Rate limiting applies cooldown windows to prevent notification spam during rapid update bursts</li>
