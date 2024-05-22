# Fairness Engine

Fairness Engine is a secure multi-party computation server that ensures input data privacy while maintaining correctness and fairness. It uses Shamir's Secret Sharing protocol for secure computation.

## Architecture

The project follows a simple client-server architecture, where clients can:

1. Generate shares of a secret using the key generation endpoint.
2. Submit their shares to the server.
3. Request the computation of the combined secret using the submitted shares.

The server handles the following operations:

1. Key Generation: Generate shares of a secret using Shamir's Secret Sharing protocol.
2. Share Submission: Collect shares submitted by clients.
3. Computation: Reconstruct the secret from the submitted shares and perform the desired computation.

## Todo

[ ] e2e test and benchmark test against other protocols
[ ] add support for multi-party async computation 
[ ] add support for resource alloc election