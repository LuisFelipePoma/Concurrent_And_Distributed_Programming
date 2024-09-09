# For ratings -> https://www.kaggle.com/datasets/kritanjalijain/amazon-reviews
# For tabular -> https://huggingface.co/datasets/polinaeterna/tabular-benchmark

from datasets import load_dataset
import pandas as pd

# Cargar el dataset
ds = load_dataset("polinaeterna/tabular-benchmark", "clf_num_Higgs")

# Guardar el dataset en disco
ds.save_to_disk('tabular-benchmark')  # type: ignore

# Convertir el dataset a DataFrame de pandas
df = ds['train'].to_pandas()

# Guardar el DataFrame como archivo CSV
df.to_csv('tabular.csv', index=False)