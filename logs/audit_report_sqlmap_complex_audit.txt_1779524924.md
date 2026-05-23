Okay, let's dissect this SQLMap output and develop a robust remediation strategy.

**Analysis & Remediation Data**

**1. Vulnerable Parameter & URL/Endpoint:**

*   **Parameter:** `target_id` (GET) – This parameter appears to be the entry point for the SQL injection attempts.
*   **URL/Endpoint:**  (Inferred)  The URL likely looks something like `/index.php?target_id=...` –  We need to examine the application's logic to determine the full endpoint.

**2. Injection Type & Explanation:**

*   **Boolean-Based Blind SQLi (First Payload):**
    *   **Type:** Boolean-based blind – SQLMap identified the presence of an `AND` clause within a `WHERE` or `HAVING` clause, typical of blind SQL injection. The response (or lack thereof) to the injected payloads allows the attacker to infer the truthiness of conditions.
    *   **Payload:** `target_id=101 AND 7312=7312`
    *   **Exploitation:** The payload attempts to test if `7312=7312` is true. If the MySQL server returns a 200 OK, the condition is true, and if it returns an error (like a 500), the condition is false. This confirms the presence of the AND clause.

*   **Error-Based SQLi (Second Payload):**
    *   **Type:** Error-based - This indicates an attempt to trigger MySQL error messages to extract information. The `GTID_SUBSET` function is frequently used to exploit the `mysql_query` function when dealing with GTID-based replication.  The `CONCAT` function is then injected to potentially overwrite the base query with the injection payload.
    *   **Payload:** `target_id=101 AND GTID_SUBSET(CONCAT(0x71767a6271,(SELECT (ELT(7312=7312,1))),0x7170707a71),7312)`
    *   **Exploitation:**  This payload injects a complex `GTID_SUBSET` function that leverages a conditional statement (`ELT`) based on the `7312=7312` boolean. If `7312=7312` is true, then the entire payload becomes  `GTID_SUBSET(...,1)`, which causes an error during query evaluation. The exact error message reveals the internal query being executed.

*   **UNION Query (Third Payload):**
    *   **Type:** Generic UNION select - The attacker attempts to retrieve data from other tables by constructing a UNION query.
    *   **Payload:** `target_id=-3815 UNION ALL SELECT NULL,CONCAT(0x71767a6271,0x546573744c6f6744617461,0x7170707a71),NULL-- -`
    *   **Exploitation:** The attacker uses `UNION ALL` to combine the result set from the original query with the data returned by the `SELECT` statement within the `UNION`. The `CONCAT` function and hex values are used to build the SQL query, likely attempting to select columns from other tables.



**3. Production-Ready, Secure Code Patch:**

The primary vulnerability here is the direct use of user-supplied input within a SQL query without proper sanitization or parameterization.

**Revised Code Snippet (PHP Example - Demonstrates Prepared Statements):**

```php
<?php
// Assume $target_id is retrieved from the GET request (e.g., $_GET['target_id'])
// Always validate and sanitize user input *before* using it in a query!

$target_id = $_GET['target_id'];
// IMPORTANT:  Validate that $target_id is an integer and within acceptable bounds.
if (!is_numeric($target_id)) {
    // Log the error and handle invalid input gracefully (e.g., return an error page).
    die("Invalid target_id provided.");
}

// Use a prepared statement to prevent SQL injection
$stmt = $pdo->prepare("SELECT * FROM users WHERE id = :target_id"); // Assuming a 'users' table with an 'id' column
$stmt->bindParam(':target_id', $target_id, PDO::PARAM_INT);
$stmt->execute();

$result = $stmt->fetch(PDO::FETCH_ASSOC);

// Process the results.
if ($result) {
  // ... Display or process user data safely ...
} else {
  // ... Handle the case where no user is found ...
}

?>
```

**Explanation:**

*   **Prepared Statements:** Prepared statements use placeholders (like `:target_id`) in the query. The database driver handles the escaping and quoting of the values, preventing them from being interpreted as SQL code.
*   **Type Casting:** `PDO::PARAM_INT` ensures that the value is treated as an integer. This helps prevent various types of injection.
*   **Input Validation:**  Explicitly check that the input is numeric before using it. This adds an extra layer of defense.
*   **Error Handling:** Implement robust error handling to catch exceptions and prevent sensitive information from being leaked.



**4. Python/Curl Confirmation Testing (PoC - Proof of Concept):**

**Python (using `requests` library):**

```python
import requests

target_url = "your_application_url/index.php?target_id="  # Replace with actual URL

# Boolean-based test
payload = "101 AND 7312=7312"
response = requests.get(target_url + payload)
print(f"Boolean Response: {response.text}")

# Error-based test
payload = "101 AND GTID_SUBSET(CONCAT(0x71767a6271,(SELECT (ELT(7312=7312,1))),0x7170707a71),7312)"
response = requests.get(target_url + payload)
print(f"Error Response: {response.text}")

#UNION Test
payload = "-3815 UNION ALL SELECT NULL,CONCAT(0x71767a6271,0x546573744c6f6744617461,0x7170707a71),NULL-- -"
response = requests.get(target_url + payload)
print(f"Union Response: {response.text}")
```

**Curl Command:**

```bash
curl -g -d "target_id=101 AND 7312=7312" "your_application_url/index.php"
curl -g -d "target_id=101 AND GTID_SUBSET(CONCAT(0x71767a6271,(SELECT (ELT(7312=7312,1))),0x7170707a71),7312)" "your_application_url/index.php"
curl -g -d "target_id=-3815 UNION ALL SELECT NULL,CONCAT(0x71767a6271,0x546573744c6f6744617461,0x7170707a71),NULL-- -" "your_application_url/index.php"
```

**Important Considerations:**

*   **Replace `your_application_url`** with the actual URL of the vulnerable application.
*   The response text from these tests will depend on the application's exact implementation. You should expect to see errors or potentially sensitive data if the vulnerability is present.
*   **Testing in a Staging Environment:** *Never* perform these tests on a production system. Always use a test or staging environment to verify the effectiveness of your patches.
*   **Defense in Depth:**  Implement multiple layers of security, including input validation, prepared statements, web application firewalls (WAFs), and regular security audits.

This detailed analysis and remediation plan will help you secure your application against these SQL injection vulnerabilities.  Remember to prioritize the use of prepared statements and robust input validation for all user-supplied data.