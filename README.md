Zero Trust Prototype

This project is a simple implementation of certain elements commonly found in Zero Trust models, such as dynamic
identity issuance and multi-factor authentication (MFA). It consists of three Go servers:

1) Dynamic Identity Server: Issues and revokes dynamic identities.
2) MFA Server: Issues MFA tokens.
3) API Server: Provides the current local time, accessible only with a valid identity and MFA token.

Prerequisites

1) Go 1.x
2) Bash (for the start-up script)

Getting Started

1. Clone the repository

```bash
git clone [Your Repository URL]
cd [Your Repository Directory]
```

2. Build and Start Servers

To build and start all three servers at once:

```bash
./start_servers.sh
```

The servers will start in the background. Their respective PIDs will be displayed, which can be useful for management
purposes.
Testing

You can test the functionality using curl or any other HTTP client. Here are a few example curl commands:

##### Obtain a dynamic identity:

```bash
curl http://localhost:8080/issue
```

... (other curl commands for MFA and API Server) ...
Implementation Details

The implementation relies on in-memory databases for both identity and MFA token storage, making it suitable for
demonstration and educational purposes but not for production use.

#### Stopping the Servers

Simply terminate the start_servers.sh script, and it will automatically stop all three servers.

#### Future Improvements

1) Implement continuous verification for identities.
2) Integrate with actual databases for persistence.
3) Add more granular access controls and policies.
4) Extend with user-based authentication and more comprehensive logging.

#### License

This project is open-source and available under the MIT license.
