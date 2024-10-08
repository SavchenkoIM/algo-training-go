# L. Пересечение отрезков

Вам даны две последовательности отрезков. Каждый отрезок задаётся координатой начала и конца — [start<sub>i</sub>,end<sub>i</sub>]. Отрезки каждой последовательности отсортированы слева направо и не имеют общих точек.

Найдите пересечение двух последовательностей отрезков. То есть третью последовательность отрезков, такую, что:

* Каждый отрезок содержится в некотором отрезке и первой, и второй последовательности;
* Никакой отрезок нельзя увеличить; 
* Отрезки этой последовательности не имеют общих точек;
* Отрезки в последовательности также отсортированы в порядке возрастания.


## Формат ввода

В первой строке дано число отрезков в первой последовательности n (0≤n≤100000)

В следующих n строках даны отрезки первой последовательности по одному на строке. Каждый отрезок записан в формате startiendi, координаты начала и конца целые неотрицательные числа, не превосходящие по модулю 109.

В строке n+2 дана длина второй последовательности m, (0≤m≤100000).

В следующих m строках заданы отрезки второй последовательности.

Гарантируется, что end<sub>i</sub><start<sub>i+1</sub>, а также что end<sub>i</sub>−start<sub>i</sub>>0.

## Формат вывода

Выведите по одному в строке отсортированные отрезки из пересечения последовательностей в том же формате, что и во входных данных. Заметьте, что длина отрезков в пересечении может быть нулевой, в этом случае start<sub>i</sub>=end<sub>i</sub>.

## Пример 1
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>3<br>
1 4<br>
5 10<br>
15 16<br>
2<br>
0 2<br>
4 5
</td><td>1 2<br>
4 4<br>
5 5
</td></tr>
</table>

## Пример 2
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>1<br>
1 4<br>
1<br>
1 4
</td><td>1 4</td></tr>
</table>


---
# Решение

Создаём два слайса отрезков, представленных структурами с полями "начало/конец", работаем в цикле до тех пор, пока какой-либо из этих слайсов не станет пустым. На каждом цикле:

Берём из обоих слайсов следующий отрезок (нулевой элемент). Принимаем за "первый" тот, у которого меньшее из двух (или равно) начало, за "второй" - другой.

Если конец "первого" меньше начала "второго" (отрезки не имеют общих точек), исключаем из соответствующего слайса "первый" отрезок, `continue`

Печатаем максимальное среди двух отрезков начало и минимальный конец. (часть ответа на задачу)

Если конец любого из отрезков меньше или равен "минимальному концу" сради двух отрезков, то такой отрезок удаляется из своего слайса (на каждом шаге будет как минимум одно удаление, за счёт чего цикл гарантированно завершится).