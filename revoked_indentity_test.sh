#!/bin/bash

# 1. Obtain a dynamic identity
DYNAMIC_IDENTITY=$(curl -s http://localhost:8080/issue | awk '{print $3}')
echo "Obtained Dynamic Identity: $DYNAMIC_IDENTITY"

# 2. Revoke the obtained dynamic identity
curl -s "http://localhost:8080/revoke?identity=$DYNAMIC_IDENTITY"
echo "Revoked Dynamic Identity: $DYNAMIC_IDENTITY"

# 3. Obtain an MFA token (for completeness)
MFA_TOKEN=$(curl -s http://localhost:8081/issue-token | awk '{print $4}')

# 4. Try to use the revoked dynamic identity and MFA token to fetch current local time
RESPONSE=$(curl -s "http://localhost:8082/current-time?identity=$DYNAMIC_IDENTITY&token=$MFA_TOKEN")

echo "API Server Response: $RESPONSE"
