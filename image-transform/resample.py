# Imports
import rasterio
from rasterio.enums import Resampling
from rasterio.plot import show, show_hist
import numpy as np
from datetime import datetime
from glob import glob
import math
import os

# Config
inputFolder = 'new'
outputFolder = 'new_resample'
scale_factor = 5/8


# timestamp = datetime.now().strftime("%Y%m%d%H%M%S") # get current timestamp
count = 0 # image names lol

if not os.path.exists(os.path.join(os.getcwd(), inputFolder)):
    os.mkdir(os.path.join(os.getcwd(), inputFolder))
if not os.path.exists(os.path.join(os.getcwd(), outputFolder)):
    os.mkdir(os.path.join(os.getcwd(), outputFolder))
    
    
    
for filename in glob(os.path.join(inputFolder, '*.tif')):
    print(filename) # Print filenames for debugging
    with rasterio.open(os.path.join(os.getcwd(), filename), 'r') as dataset:
        # resample data to target shape
        data = dataset.read(
            out_shape=(
                dataset.count,
                int(dataset.height * scale_factor),
                int(dataset.width * scale_factor)
            ),
            resampling=Resampling.nearest # not using bilinear/cubic/some other shit since we dont have continuous data
        )

        # scale image transform
        transform = dataset.transform * dataset.transform.scale(
            (dataset.width / data.shape[-1]),
            (dataset.height / data.shape[-2])
        )
        
        profile = dataset.profile
        profile.update({
            'height' : int(dataset.height * scale_factor),
            'width': int(dataset.width * scale_factor),
            'transform': transform
        })
    
        with rasterio.open(os.path.join(os.getcwd(), outputFolder, str(count)+'.tif'), 'w', **profile) as dst: 
            dst.write(dataset.read())

    count += 1
