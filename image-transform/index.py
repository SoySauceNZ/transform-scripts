from rasterio.windows import Window
from datetime import datetime
from glob import glob
import rasterio
import math
import os

# Config
size = 500  # Current square size of the image
csvFolder = 'csv'
inputFolder = 'new_resample'
outputFolder = 'new_output'

# Base csv on current timestamp
timestamp = datetime.now().strftime("%Y%m%d%H%M%S")

count = 0

if not os.path.exists(os.path.join(os.getcwd(), csvFolder)):
    os.mkdir(os.path.join(os.getcwd(), csvFolder))
if not os.path.exists(os.path.join(os.getcwd(), inputFolder)):
    os.mkdir(os.path.join(os.getcwd(), inputFolder))
if not os.path.exists(os.path.join(os.getcwd(), outputFolder)):
    os.mkdir(os.path.join(os.getcwd(), outputFolder))

# Create csv file for adding LatLng
with open(os.path.join(os.getcwd(), csvFolder, timestamp+'.csv'), 'wt+') as append:
    os.mkdir(os.path.join(os.getcwd(), outputFolder, timestamp))
    append.write('image,filename,lattop,lngleft,latbottom,lngright\n')  # Header of the csv

    # Find all images in the ./image/*.tif folder
    for filename in glob(os.path.join(inputFolder, '*.tif')):
        print(filename)  # Print filenames for debugging

        # Open .tif file using Rasterio
        with rasterio.open(os.path.join(os.getcwd(), filename), 'r') as raster:
            bound = raster.bounds
            height = raster.height
            width = raster.width
            raster.close()

            # Move horizontally then vertically
            for x in range(0, width, size):
                for y in range(0, height, size):
                    if ((y+size > height) or (x+size > width)):
                        continue

                    # Calculate LatLng using ratios of the image, Ratio should not match LatLng
                    # as lat lng is designed for globe coords. Using ratio as is much easier to
                    # calculate. Also our data for LatLng isn't very accurate so the difference
                    # should not impact our results.
                    window = Window(x, y, size, size)

                    with rasterio.open(os.path.join(os.getcwd(), filename)) as raster2:
                        transform = raster2.window_transform(window)

                        profile = raster2.profile
                        profile.update({
                            'height': size,
                            'width': size,
                            'transform': transform
                        })

                        data = raster2.read(window=window)

                        with rasterio.open(os.path.join(os.getcwd(), outputFolder, timestamp, str(count)+'.tif'), 'w', **profile) as dst:
                            dst.write(data)
                    
                    b = None
                    with rasterio.open(os.path.join(os.getcwd(), outputFolder, timestamp, str(count)+'.tif'), 'r') as raster:
                        print(raster.bounds)
                        b = raster.bounds
                    
                    os.rename(os.path.join(os.getcwd(), outputFolder, timestamp, str(count)+'.tif'),os.path.join(os.getcwd(), outputFolder, timestamp, str(b.top) + "_" + str(b.left)+'.tif'))

                    # append to count of the image in relation to csv and LatLng coords
                    append.write(str(b.top) + "_" + str(b.left)+','+filename+','+str(b.top)+',' +
                                 str(b.left)+','+str(b.bottom)+','+str(b.right)+'\n')
                    count += 1
