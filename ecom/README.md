# Designing API in golang

```
#Architecure
 -------
| product|          Orders              Order Items
|--------|          _______            ------------- 
| id     |          id                  order_id
| name  |     -->   customer_id  -->    quantitiy 
|price  |           created_at          price
|quantity|          status              product_id
 --------
```

[video link](https://www.youtube.com/watch?v=s3XItrqfccw)
