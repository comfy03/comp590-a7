# comp590-a7

# Assignment 7: Sleeping Barber Problem

## Team Members:
- Comfort Donkor

## Description:
This Go program simulates the classic **Sleeping Barber Problem**, a synchronization problem in computer science. The program demonstrates the interaction between barbers and customers in a barber shop with limited waiting room capacity. It uses goroutines and channels to simulate concurrency, along with mutexes for synchronization.

### Key Features:
1. **Barber Behavior**:
   - Barbers sleep if there are no customers in the waiting room.
   - Barbers serve customers in the order they arrive.
   - Supports multiple barbers working simultaneously.

2. **Customer Behavior**:
   - Customers either take a seat in the waiting room or leave if the room is full.
   - Each customer is either served or departs based on waiting room availability.

3. **Customizability**:
   - Users can specify the number of chairs in the waiting room, the number of customers, the time taken by barbers to cut hair, and the number of barbers.

