
# Hadahoop Distributed File Storage System

## Project Overview

Hadahoop is my take on building a scalable and resilient file storage system using distributed principles. By spreading data across a network of nodes, it ensures redundancy, faster access, and robustness against failures, making it perfect for modern applications that require large-scale data storage. I built it as a personal project to dive deeper into distributed systems and explore the practical applications of technologies like Hadoop, IPFS, and cloud storage patterns.

### Features

- **Distributed File Storage:** Stores files in chunks across multiple nodes to improve scalability and fault tolerance.
- **Data Redundancy:** Ensures data availability by replicating files across nodes to prevent data loss.
- **Fault Tolerant:** If a node goes down, the system can still access the data by pulling it from other nodes.
- **High Availability:** Multiple nodes ensure the system stays online even when some parts of the network fail.
- **Fast File Retrieval:** Optimized for quick file access, with a distributed hash table for efficient searching.

### Technologies Used

- **Go:** The backbone language for implementing the core features, node communication, and file management.
- **Distributed Systems:** Concepts and design inspired by Hadoop's distributed file system and IPFS.
- **SQLite:** Used for lightweight local storage for metadata management (could be swapped with a more robust database for production use).

### How It Works

Hadahoop breaks large files into smaller chunks and stores these chunks on multiple nodes within the system. Each chunk is replicated for redundancy, and metadata about the chunks (e.g., where they’re stored) is maintained in a central index. The API provides endpoints to upload files, retrieve them, and delete them.

- **Uploading Files:** When a file is uploaded, it's split into chunks. Each chunk is stored on a different node, and the metadata for the file is saved.
- **Retrieving Files:** To retrieve a file, the system finds the chunks from the nodes based on the stored metadata and assembles the file.
- **Fault Tolerance:** If a node goes down, the system automatically retrieves the file from its replicated chunks stored on other nodes.

### My Experience

Building this system was a lot of fun and challenging at the same time. I learned a lot about distributed computing, fault tolerance, and scalability. Implementing the file chunking and replication system made me realize how important redundancy is for ensuring data safety. Plus, working with Docker and creating a REST API really brought everything together for seamless deployment and interaction.

The best part? I can easily scale it up by adding more nodes, which is essential for real-world applications. I’m excited about the potential to expand this project further — maybe integrating encryption for added security or exploring a hybrid cloud approach.

### Future Improvements

- **Security Enhancements:** Adding encryption for files during storage and transmission.
- **Data Compression:** Implementing a compression mechanism to reduce storage requirements.
- **Advanced API Features:** Adding more advanced search and filtering capabilities.
- **Cloud Integration:** Integrating the system with cloud platforms like AWS or Azure for hybrid storage.

Feel free to reach out if you have any questions or want to contribute! 😊
