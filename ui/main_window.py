# -*- coding: utf-8 -*-

# Form implementation generated from reading ui file '.\ui_templates\MainWindow.ui'
#
# Created by: PyQt5 UI code generator 5.15.9
#
# WARNING: Any manual changes made to this file will be lost when pyuic5 is
# run again.  Do not edit this file unless you know what you are doing.


from PyQt5 import QtCore, QtGui, QtWidgets
import sys


class Ui_MainWindow(object):

    def setupUi(self, MainWindow):
        MainWindow.setObjectName("MainWindow")
        MainWindow.resize(640, 476)
        MainWindow.setMinimumSize(QtCore.QSize(640, 476))
        MainWindow.setMaximumSize(QtCore.QSize(640, 476))
        self.centralwidget = QtWidgets.QWidget(MainWindow)
        self.centralwidget.setObjectName("centralwidget")
        self.verticalLayoutWidget_3 = QtWidgets.QWidget(self.centralwidget)
        self.verticalLayoutWidget_3.setGeometry(QtCore.QRect(10, 10, 621, 431))
        self.verticalLayoutWidget_3.setObjectName("verticalLayoutWidget_3")
        self.verticalLayout_5 = QtWidgets.QVBoxLayout(self.verticalLayoutWidget_3)
        self.verticalLayout_5.setContentsMargins(0, 0, 0, 0)
        self.verticalLayout_5.setObjectName("verticalLayout_5")
        self.horizontalLayout = QtWidgets.QHBoxLayout()
        self.horizontalLayout.setObjectName("horizontalLayout")
        self.textEdit = QtWidgets.QTextEdit(self.verticalLayoutWidget_3)
        self.textEdit.setEnabled(True)
        self.textEdit.setMaximumSize(QtCore.QSize(501, 24))
        self.textEdit.setReadOnly(True)
        self.textEdit.setObjectName("textEdit")
        self.horizontalLayout.addWidget(self.textEdit)
        self.pushButton = QtWidgets.QPushButton(self.verticalLayoutWidget_3)
        self.pushButton.setObjectName("pushButton")
        self.horizontalLayout.addWidget(self.pushButton)
        self.verticalLayout_5.addLayout(self.horizontalLayout)
        self.horizontalLayout_2 = QtWidgets.QHBoxLayout()
        self.horizontalLayout_2.setObjectName("horizontalLayout_2")
        self.label = QtWidgets.QLabel(self.verticalLayoutWidget_3)
        self.label.setObjectName("label")
        self.horizontalLayout_2.addWidget(self.label)
        self.radioButton = QtWidgets.QRadioButton(self.verticalLayoutWidget_3)
        self.radioButton.setChecked(True)
        self.radioButton.setObjectName("radioButton")
        self.horizontalLayout_2.addWidget(self.radioButton)
        self.radioButton_2 = QtWidgets.QRadioButton(self.verticalLayoutWidget_3)
        self.radioButton_2.setObjectName("radioButton_2")
        self.horizontalLayout_2.addWidget(self.radioButton_2)
        self.verticalLayout_5.addLayout(self.horizontalLayout_2)
        self.pushButton_2 = QtWidgets.QPushButton(self.verticalLayoutWidget_3)
        self.pushButton_2.setObjectName("pushButton_2")
        self.verticalLayout_5.addWidget(self.pushButton_2)
        self.textEdit_2 = QtWidgets.QPlainTextEdit(self.verticalLayoutWidget_3)
        self.textEdit_2.setReadOnly(True)
        self.textEdit_2.setObjectName("textEdit_2")
        self.verticalLayout_5.addWidget(self.textEdit_2)
        self.verticalLayout = QtWidgets.QVBoxLayout()
        self.verticalLayout.setObjectName("verticalLayout")
        self.horizontalLayout_3 = QtWidgets.QHBoxLayout()
        self.horizontalLayout_3.setObjectName("horizontalLayout_3")
        self.label_2 = QtWidgets.QLabel(self.verticalLayoutWidget_3)
        self.label_2.setObjectName("label_2")
        self.horizontalLayout_3.addWidget(self.label_2)
        self.verticalLayout_2 = QtWidgets.QVBoxLayout()
        self.verticalLayout_2.setObjectName("verticalLayout_2")
        self.pushButton_4 = QtWidgets.QPushButton(self.verticalLayoutWidget_3)
        self.pushButton_4.setObjectName("pushButton_4")
        self.verticalLayout_2.addWidget(self.pushButton_4)
        self.pushButton_3 = QtWidgets.QPushButton(self.verticalLayoutWidget_3)
        self.pushButton_3.setObjectName("pushButton_3")
        self.verticalLayout_2.addWidget(self.pushButton_3)
        self.pushButton_5 = QtWidgets.QPushButton(self.verticalLayoutWidget_3)
        self.pushButton_5.setObjectName("pushButton_5")
        self.verticalLayout_2.addWidget(self.pushButton_5)
        self.horizontalLayout_3.addLayout(self.verticalLayout_2)
        self.verticalLayout.addLayout(self.horizontalLayout_3)
        self.verticalLayout_5.addLayout(self.verticalLayout)
        MainWindow.setCentralWidget(self.centralwidget)
        self.statusbar = QtWidgets.QStatusBar(MainWindow)
        self.statusbar.setObjectName("statusbar")
        MainWindow.setStatusBar(self.statusbar)

        self.retranslateUi(MainWindow)
        QtCore.QMetaObject.connectSlotsByName(MainWindow)

    def retranslateUi(self, MainWindow):
        _translate = QtCore.QCoreApplication.translate
        MainWindow.setWindowTitle(_translate("MainWindow", "Antivirus"))
        self.pushButton.setText(_translate("MainWindow", "Выбрать..."))
        self.label.setText(_translate("MainWindow", "Вид сканирования:"))
        self.radioButton.setText(_translate("MainWindow", "Мониторинг"))
        self.radioButton_2.setText(_translate("MainWindow", "Сканирование по расписанию"))
        self.pushButton_2.setText(_translate("MainWindow", "Запуск сканирования"))
        self.label_2.setText(_translate("MainWindow", ""))
        self.pushButton_4.setText(_translate("MainWindow", "Очистить окно вывода"))
        self.pushButton_3.setText(_translate("MainWindow", "Очистить карантин"))
        self.pushButton_5.setText(_translate("MainWindow", "Остановить сервис"))



