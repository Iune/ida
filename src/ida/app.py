"""
Process voters for music competitions
"""
import sys
from PySide2 import QtWidgets


class Ida(QtWidgets.QMainWindow):
    def __init__(self):
        super().__init__()
        self.init_ui()

    def init_ui(self):
        self.setWindowTitle('ida')
        self.show()

def main():
    app = QtWidgets.QApplication(sys.argv)
    main_window = Ida()
    sys.exit(app.exec_())
