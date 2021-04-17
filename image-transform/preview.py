# Imports
import rasterio
import numpy as np
import os
from rasterio.warp import reproject, Resampling

filename0 = "new/latlng.tif"
filename0="C:/Users/cheng/Desktop/image-transform/new_output/20210417142300/-36.87738742212947_174.817098082761.tif"
dst_crs = 'EPSG:4326'

with rasterio.open(os.path.join(os.getcwd(), filename0), 'r') as raster:
    print(raster.bounds)
    bound = raster.bounds
    new_bounds = rasterio.warp.transform_bounds(raster.crs, dst_crs, bound[0], bound[1], bound[2], bound[3])
    print(new_bounds)
