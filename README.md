## Csv to geojeson converting tool

This application makes you possible to convert rapidly geo-data visualization. Then, You will be prompted a csv file having headers.

Supported geometry type

- Point
- LineString
- Polygon

Not supported mutil geometry type. That is feature work.

Example: Converting csv file to geojson file

```
csv2geojson -t Point -lon longitude -lat latitude ./path/example.csv
```

Example: Converting tsv file to geojson file

```
csv2geojson -d $'\t' -t LineString -lon longitude -lat latitude ./path/example.csv
```
