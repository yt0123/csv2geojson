## Csv to geojeson converting tool

This application makes you possible to convert rapidly geo-data visualization. Then, You will be prompted a csv file having headers.

Usage:

&nbsp; csv2geojson [OPTIONS] PATTERN [PATH]

Application Options:

&nbsp; -o, --output=    Set output path for converted geojson file

&nbsp; -t, --type=      Set geometry type for geojson file (default: Point)

&nbsp; -k, --key=       Set key column to join records
  -d, --delimiter= Set csv delimiter for imported csv file (default: ,)

&nbsp;&nbsp; --lon=       Set geometry coordinates for geojson file (default: longitude)

&nbsp;&nbsp; --lat=       Set geometry coordinates for geojson file (default: latitude)

&nbsp;&nbsp; --quotes     Check csv double quotes for imported csv file

&nbsp; -p, --preformat  Output preformatted geojson file
&nbsp; -v, --verbose    Show verbose debug information

Help Options:

&nbsp; -h, --help       Show this help message

Supported geometry type

- Point
- LineString
- Polygon

Not supported mutil geometry type. That is feature work.


**Examples**

Converting csv file to geojson file

```
csv2geojson -t Point -lon longitude -lat latitude ./path/example.csv
```

Converting tsv file to geojson file

```
csv2geojson -d $'\t' -t LineString -lon longitude -lat latitude ./path/example.csv
```
