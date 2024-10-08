# B. Card Counter

На стол в ряд выложены карточки, на каждой карточке написано натуральное число. За один ход разрешается взять карточку либо с левого, либо с правого конца ряда. Всего можно сделать k ходов. Итоговый счет равен сумме чисел на выбранных карточках. Определите, какой максимальный счет можно получить по итогам игры.  

## Формат ввода

В первой строке записано число карточек n (1≤n≤10<sup>5</sup>).

Во второй строке записано число ходов k (1≤k≤n).

В третьей строке через пробел даны числа, записанные на карточках. i-е по счету число записано на i-й слева карточке. Все числа натуральные и не превосходят 10<sup>4</sup>.

## Формат вывода

Выведите единственное число —- максимальную сумму очков, которую можно набрать, сделав k ходов.
## Пример 1
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>7<br>
3<br>
5 8 2 1 3 4 11</td><td>24</td></tr>
</table>

## Пример 2
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>5<br>
5<br>
1 2 3 4 5</td><td>15</td></tr>
</table>

## Пример 3
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>7<br>
4<br>
1 1 9 2 2 2 6</td><td>17</td></tr>
</table>


---
# Решение

Здесь главное не пытаться просчитывать всё дерево возможных комбинаций порядка взятия карт.  

Для нас важно только то, сколько карт возьмут справа, а сколько слева. Для примера 3 возможно 5 (k+1) комбинаций:  

* все слева: 1 1 9 2
* 3 слева, 1 справа: 1 1 9, 6
* по 2 с каждой стороны: 1 1, 2 6
* одна слева, 3 справа: 1, 2 2 6
* все справа: 2 2 2 6

Лучше всего вторая: 1 1 9, 6 — ответ 17