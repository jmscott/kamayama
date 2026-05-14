#
#  Synopsis:
#	Makefile for www development environment of $(KAMAYAMA_ROOT)
#  Note:
#	Move "ifdef NOC_DNS_VHOST_SUFFIX" into noc/Makefile
#
include ../local.mk

WWW_ROOT=$(KAMAYAMA_PREFIX)/www

all:

install:
	install -d $(KAMAYAMA_PREFIX)
	install -d $(WWW_ROOT)
	install -m u=rwx,g=rx,o= -d $(WWW_ROOT)/lib
	install -m u=rwx,g=rx,o= -d $(WWW_ROOT)/cgi-bin
	install -m ug=wrx,o= -d $(WWW_ROOT)/tmp
	install -m u=rwx,g=rx,o= -d $(WWW_ROOT)/htdocs
	install -m u=rw,g=r,o= index.html $(WWW_ROOT)/htdocs

distclean:
	rm -rf $(WWW_ROOT)/cgi-bin
	rm -rf $(WWW_ROOT)/htdocs
	rm -rf $(WWW_ROOT)/lib

clean:

world:
	$(MAKE) $(MFLAGS) clean
	$(MAKE) $(MFLAGS) all
	$(MAKE) $(MFLAGS) distclean
	$(MAKE) $(MFLAGS) install
tar:
	make-make tar kamayama-www.mkmk
