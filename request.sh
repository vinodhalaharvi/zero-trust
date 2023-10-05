# 1. Obtain a dynamic identity
DYNAMIC_IDENTITY=$(curl -s http://localhost:8080/issue | awk '{print $3}')

# 2. Obtain an MFA token
MFA_TOKEN=$(curl -s http://localhost:8081/issue-token | awk '{print $4}')

# 3. Use both the dynamic identity and MFA token to fetch current local time
curl "http://localhost:8082/current-time?identity=$DYNAMIC_IDENTITY&token=$MFA_TOKEN"

