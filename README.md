Customer Behavior Analytics Tool

Overview

The Customer Behavior Analytics Tool is designed to process a stream of customer activity events from a JSONL file. The application analyzes customer behavior on an e-commerce platform, identifying trends such as popular products and sales performance. The tool focuses on real-time and historical insights over predefined time frames, including metrics like the most viewed item, conversion rate from views to purchases, and highest-grossing item.

Functionality

1. Event Processing:
    - The tool processes events like page_visit, add_to_cart, and purchase.
    - Each event contains attributes like customer_id, timestamp, and item_id.

2. Metrics Calculation:
    - Most Viewed Item: The item with the highest number of page views.
    - Conversion Rate: The percentage of views that converted into cart additions or purchases.
    - Highest-Grossing Item: The item with the highest revenue.

3. Time Frames:
    - The tool considers various time frames, such as the last hour or the last 24 hours.
    - Events are aggregated into hourly and daily buckets.

4. Data Aging:
    - Data is aged and aggregated into broader time buckets using in-memory structures.

Requirements

- Go: The tool is written in Go. Ensure you have Go installed on your machine. You can download it from the official Go website (https://golang.org/).

Running the Application

1. Clone the Repository:

   git clone <repository-url>
   cd <repository-directory>

2. Prepare the Input:

   - Place the input JSONL file (sample_events.jsonl) in the same directory as the Go application or adjust the file path in the code.

3. Run the application
   `go run .`

   - The application will read the events from sample_events.jsonl and output the aggregate statistics.

Example Output

Time Bucket: 2024-03-27 05:00:00 +0000 UTC
Conversion rates:
Item ABC: 10.00%
Most viewed item: XYZ with 20 views
Highest-grossing item: PQR with 50.00 revenue

Design Decisions

1. Event Handling:
   - Events are represented by the Event struct, which holds common event attributes.

2. Bucket Structure:
   - Events are grouped into Bucket structs representing specific time frames.
   - Each bucket aggregates metrics for its set of events.

3. In-Memory Storage:
   - The tool uses in-memory maps and slices to store and process event data.
   - This approach is efficient for the scope of the assignment.

4. Scaling:
   - For production use, the tool could be scaled by offloading data to a database or a data lake, and using a streaming platform for distributed processing.

Future Improvements

1. Database Integration:
   - The tool can be integrated with a database for persistent storage and querying.
   - Time-series Databases (ex. InfluxDB): Ideal for storing and querying data with timestamps as they offer efficient data ingestion and retrieval for time-based queries (e.g., hourly or daily aggregates).

2. Distributed Processing:
   - The tool can leverage a streaming platform like Kafka for distributed event processing.

3. Enhanced Insights:
   - Additional insights, such as customer retention or product affinity, can be added based on business needs.

Conclusion

The Customer Behavior Analytics Tool is a simple yet scalable system for processing and analyzing customer activity events on an e-commerce platform. The tool demonstrates knowledge of data structures, algorithms, and scalable system design.
