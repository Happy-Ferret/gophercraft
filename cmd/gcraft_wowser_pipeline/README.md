# gcraft_wowser_pipeline

Attempting to make a game content server for Wowser, and for serving world tiles

# Architecture

```             
        blizzardry/mpq      various format decoders
+-----------+      +------------------------+
| game MPQs |______| gcraft_wowser_pipeline |
+-----------+      +------------------------+
                    |     |         |     
                    |     | HTTP/2  |  Alternate protocol?
                    |     |         | 
                    |     |         |
                    |   +-----------------------+         (optional)           Map data +-------------------+
                    |   | gcraft_pipeline_cache |---------------------------------------| gcraft_core_world |
            +-------+   +-----------------------+                                       +-------------------+
            |           | Gzip compresses and stores game assets in a folder
            |           | Avoid extracting same MPQ file repeatedly.
            |   HTTP/2  | Provide same API as Wowser's StormLib-based NodeJS pipeline.
            |           | Control rate limiting, etc.
            +-----------------------+
            | Wowser client         |
            +-----------------------+
```