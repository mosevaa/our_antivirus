import subprocess
import sys, os, time, threading
from main_window import *
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QFileDialog
from subprocess import Popen, PIPE

os.environ['PYTHONIOENCODING'] = 'utf-8'


class MainWin(QtWidgets.QMainWindow):
    def __init__(self):
        QtWidgets.QWidget.__init__(self)
        self.ui = Ui_MainWindow()
        self.ui.setupUi(self)
        self.ui.pushButton_2.clicked.connect(self.start)

    def message(self, m):
        self.ui.textEdit_2.appendPlainText(m)

    def start(self):
        self.message('start')
        cmd = 'go run C:\\Users\\Lesya\\GolandProjects\\our_antivirus\\av\\main.go'
        p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        stdout, stderr = p.communicate()
        self.ui.textEdit_2.appendPlainText(str(stdout.decode()+stderr.decode()))


if __name__ == "__main__":
    app = QtWidgets.QApplication(sys.argv)
    MyApp = MainWin()
    MyApp.show()
    sys.exit(app.exec_())
