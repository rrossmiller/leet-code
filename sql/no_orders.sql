SELECT name AS Customers FROM orders AS o
RIGHT JOIN customers AS c
ON o.customerId = c.id
WHERE o.customerId IS NULL;