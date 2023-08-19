CC = gcc
CFLAGS = -I/usr/include/mysql -g
LDFLAGS = -L/usr/lib/mysql -lmysqlclient

SOURCES = login.c
EXECUTABLE = loginSecurithLib

all: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)

clean:
	rm -f $(EXECUTABLE)


