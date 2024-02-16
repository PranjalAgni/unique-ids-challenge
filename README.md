# Unique ID generator 🧵

In this challenge, you’ll need to implement a globally-unique ID generation system that runs against Maelstrom’s unique-ids workload. Your service should be totally available, meaning that it can continue to operate even in the face of network partitions.

## Approaches I used

1. Using a uuid prefix for each node which is generated when the node boot's up, and then for each request which comes to the node we generate id like = prefix + time_in_nanoseconds

2. Implementing a UUID4 algorithm from scratch, this will make sure it's golbally unique (8-4-4-12)
