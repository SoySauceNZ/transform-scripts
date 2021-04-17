# Image-Transform

Used to transform images from large .tif files into smaller images based on LatLng of the tif.

Currently used to cut tif into smaller square images 500x500. The cutting starts from (0, 0) 
pixels and cut based on defined `size` 

The transformer calculates LatLng using ratios of based on the crop of the image, ratio should 
not match LatLng as lat lng is designed for globe coords. Using ratio as is much easier to 
calculate. Also our data for LatLng isn't very accurate so the difference should not impact 
our results.

## Install

Install [Rasterio](https://rasterio.readthedocs.io/en/latest/installation.html)

Installing Rasterio depends on platform, example for windows users have to download built binaries.

## Setup

Place all `*.tif` files into `./image` folder the transformer would go through files using


## Run

```bash
python index.py
```