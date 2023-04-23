import sys, os, time, threading
from main_window import *
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QFileDialog
from subprocess import Popen, PIPE

os.environ['PYTHONIOENCODING'] = 'utf-8'


def buffer_to_str(buf):
    codec = QtCore.QTextCodec.codecForName("UTF-8")
    return str(codec.toUnicode(buf))


class Process(QtCore.QObject):
    def __init__(self):
        self.stdout = QtCore.pyqtSignal(str)
        self.stderr = QtCore.pyqtSignal(str)
        self.finished = QtCore.pyqtSignal(bool)

    def start(self, cmd, args):
        process = QtCore.QProcess()
        process.setProgram(cmd)
        process.setArguments(args)
        process.readyReadStandardError.connect(lambda: self.stderr.emit(buffer_to_str(process.readAllStandardError())))
        process.readyReadStandardOutput.connect(lambda: self.stderr.emit(buffer_to_str(process.readAllStandardOutput())))
        process.finished.connect(self.finished)
        process.start()

        self._process = process
