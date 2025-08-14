QT       += core gui widgets sql

CONFIG   += c++17

SOURCES  += main.cpp \
            src/mainwindow.cpp \
            src/database.cpp

HEADERS  += include/mainwindow.h \
            include/database.h

OBJECTS_DIR = build
MOC_DIR = build
RCC_DIR = build
UI_DIR = build
