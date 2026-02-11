# Ответы на задания семинара 6

## Часть 1: Select

### Вопрос 1
Показать все названия книг вместе с именами издателей.

```sql
SELECT title, publisher_name
FROM book;
```

### Вопрос 2
В какой книге наибольшее количество страниц?

```sql
SELECT title, author, number_of_pages
FROM book
WHERE number_of_pages = (
    SELECT MAX(number_of_pages)
    FROM book
);
```

### Вопрос 3
Какие авторы написали более 5 книг?

```sql
SELECT author, COUNT(isbn) AS books_count
FROM book
GROUP BY author
HAVING COUNT(isbn) > 5
ORDER BY books_count DESC;
```

### Вопрос 4
В каких книгах более чем в два раза больше страниц, чем среднее количество страниц для всех книг?

```sql
SELECT title, author, number_of_pages
FROM book
WHERE number_of_pages > 2 * (
    SELECT AVG(number_of_pages)
    FROM book
);
```

### Вопрос 5
Какие категории содержат подкатегории?

```sql
SELECT DISTINCT parentcat
FROM category
WHERE parentcat IS NOT NULL;
```

### Вопрос 6
У какого автора (предположим, что имена авторов уникальны) написано максимальное количество книг?

```sql
SELECT author, COUNT(isbn)
FROM book
GROUP BY author
HAVING COUNT(isbn) = (
    SELECT MAX(books_count)
    FROM (
        SELECT COUNT(isbn) AS books_count
        FROM book
        GROUP BY author
    ) AS t
);
```

### Вопрос 7
Какие читатели забронировали все книги (не копии), написанные "Марком Твеном"?

```sql
SELECT r.firstname, r.lastname
FROM reader r
WHERE NOT EXISTS (
    SELECT 1
    FROM book b
    WHERE b.author = 'Марк Твен'
    AND NOT EXISTS (
        SELECT 1
        FROM borrowing br
        WHERE br.id = r.id
        AND br.isbn = b.isbn
    )
);
```

### Вопрос 8
Какие книги имеют более одной копии?

```sql
SELECT isbn, COUNT(*) AS count
FROM copy
GROUP BY isbn
HAVING COUNT(*) > 1;
```

### Вопрос 9
ТОП 10 самых старых книг

```sql
WITH BooksRank AS (
    SELECT title, author, year, DENSE_RANK() OVER (ORDER BY year) AS year_rank
    FROM book
)
SELECT title, author, year
FROM BooksRank
WHERE year_rank <= 10;
```

Если надо просто топ-10:
```sql
SELECT title, author, year
FROM book
ORDER BY year
LIMIT 10;
```

### Вопрос 10
Перечислите все категории в категории "Спорт" (с любым уровнем вложености).

```sql
WITH RECURSIVE SportCategories AS (
    SELECT categoryname
    FROM category
    WHERE categoryname = 'Спорт'

    UNION ALL

    SELECT c.categoryname AS categoryname
    FROM category c
    JOIN SportCategories sc ON c.parentcat = sc.categoryname
)
SELECT categoryname
FROM SportCategories
WHERE categoryname <> 'Спорт';
```

## Часть 2: Insert / Update / Delete

### Вопрос 1
Добавьте запись о бронировании читателем 'Василеем Петровым' книги с ISBN 123456 и номером копии 4.

```sql
INSERT INTO borrowing (id, isbn, copynumber)
SELECT id, '123456', 4
FROM reader
WHERE firstname = 'Василий' AND lastname = 'Петров'
LIMIT 1;
```

### Вопрос 2
Удалить все книги, год публикации которых превышает 2000 год.

```sql
DELETE FROM book
WHERE year > 2000;
```

### Вопрос 3
Измените дату возврата для всех книг категории "Базы данных", начиная с 01.01.2016, чтобы они были в заимствовании на 30 дней дольше (предположим, что в SQL можно добавлять числа к датам).

```sql
UPDATE borrowing b
SET returndate = b.returndate + INTERVAL '30 days'
FROM bookcategory bc
WHERE b.isbn = bc.isbn
  AND bc.categoryname = 'Базы данных'
  AND b.returndate > '01.01.2016';
```

## Часть 3: Интерпретация запросов

### Запрос 1

```sql
SELECT s.Name, s.MatrNr 
FROM Student s
WHERE NOT EXISTS (
    SELECT * 
    FROM Check c 
    WHERE c.MatrNr = s.MatrNr 
      AND c.Note >= 4.0 
);
```

**Опишите на русском языке результат запроса выше:**  
Вывести имя студента и его номер, если у него нет оценок 4.0 или выше.


### Запрос 2
```sql
( 
    SELECT p.ProfNr, p.Name, sum(lec.Credit)
    FROM Professor p, Lecture lec
    WHERE p.ProfNr = lec.ProfNr
    GROUP BY p.ProfNr, p.Name
)
UNION
( 
    SELECT p.ProfNr, p.Name, 0
    FROM Professor p
    WHERE NOT EXISTS (
        SELECT * 
        FROM Lecture lec 
        WHERE lec.ProfNr = p.ProfNr 
));
```

**Опишите на русском языке результат запроса выше:**  
Вывести номер, имя профессора и сумму кредитов за его лекции. Если у преподавателя нет лекций, заполнить сумму нулями.


### Запрос 3
```sql
SELECT s.Name, p.Note
FROM Student s, Lecture lec, Check c
WHERE s.MatrNr = c.MatrNr 
  AND lec.LectNr = c.LectNr 
  AND c.Note >= 4
  AND c.Note >= ALL (
    SELECT c1.Note 
    FROM Check c1 
    WHERE c1.MatrNr = c.MatrNr 
  )
```

**Опишите на русском языке результат запроса выше:**  
Вывести имя студента и его оценку, если оценка не меньше 4 и не меньше самой высокой оценки этого студента.
