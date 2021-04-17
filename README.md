# Transform Scripts

The following are scripts used to transform our data from CAS Open Data and Maxar satelite imagery.

## Scripts

### geojson-to-csv

Written in Go, used to convert geojson into csv format for easy read/write. Using Go allows us to parse the very large data set from CAS Open Data. This also only extract fields that we are interested in.

### image-transform

Written in Python, uses rasterio to resample, partition, preview tif images. The script also extracts lat & lng for each partition. Coords are exported to csv.

### csv-transform

Written in Go, used to merge image coords csv with CAS Open Data csv to create a dataset that links each CAS row with an associated image using coordinates and viewport.

