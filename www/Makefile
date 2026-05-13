#
#  Synopsis:
#	Makefile for www development environment of $(KAMAYAMA_ROOT)
#  Note:
#	Move "ifdef NOC_DNS_VHOST_SUFFIX" into noc/Makefile
#
include ../local.mk

WWW_ROOT=$(KAMAYAMA_PREFIX)/www

all:
ifdef NOC_DNS_VHOST_SUFFIX
	cd noc && $(MAKE) $(MFLAGS) all
endif

install:
	install -g $(KAMAYAMA_GROUP) -o $(KAMAYAMA_USER)		\
		-d $(KAMAYAMA_PREFIX)
	install -g $(KAMAYAMA_GROUP) -o $(KAMAYAMA_USER) 		\
		-d $(WWW_ROOT)
	install -g $(KAMAYAMA_GROUP) -o $(KAMAYAMA_USER) 		\
		-d $(WWW_ROOT)/lib
	install -g $(KAMAYAMA_GROUP) -o $(KAMAYAMA_USER) 		\
		-m a+wrx						\
		-d $(WWW_ROOT)/tmp
	install -g $(KAMAYAMA_GROUP) -o $(KAMAYAMA_USER) 			\
		-d $(WWW_ROOT)/htdocs

	install -g $(KAMAYAMA_GROUP) -o $(KAMAYAMA_USER) -m ug=r,o=		\
		apache2-vhost.conf.example				\
		utf82blob.pl						\
		$(WWW_ROOT)/lib
ifdef NOC_DNS_VHOST_SUFFIX
	cd noc && $(MAKE) $(MFLAGS) install
endif

distclean:
ifdef NOC_DNS_VHOST_SUFFIX
	cd noc && $(MAKE) $(MFLAGS) distclean
endif
	rm -rf $(WWW_ROOT)/htdocs
	rm -rf $(WWW_ROOT)/lib

clean:
ifdef NOC_DNS_VHOST_SUFFIX
	cd noc && $(MAKE) $(MFLAGS) clean
endif

world:
	$(MAKE) $(MFLAGS) clean
	$(MAKE) $(MFLAGS) all
	$(MAKE) $(MFLAGS) distclean
	$(MAKE) $(MFLAGS) install
tar:
	make-make tar kamayama-www.mkmk
