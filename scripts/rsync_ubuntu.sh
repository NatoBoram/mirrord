#/bin/dash

# https://wiki.ubuntu.com/Mirrors/Scripts

fatal() {
	echo "$1"
	exit 1
}

warn() {
	echo "$1"
}

# Find a source mirror near you which supports rsync on
# https://launchpad.net/ubuntu/+archivemirrors
# rsync://<iso-country-code>.rsync.archive.ubuntu.com/ubuntu should always work
RSYNCSOURCE=rsync://ca.rsync.archive.ubuntu.com/ubuntu

# Define where you want the mirror-data to be on your mirror
BASEDIR=/mnt/Seagate/mirrors/ubuntu/ubuntu

if [ ! -d ${BASEDIR} ]; then
	warn "${BASEDIR} does not exist yet, trying to create it..."
	mkdir -p ${BASEDIR} || fatal "Creation of ${BASEDIR} failed."
fi

rsync --archive --human-readable --info=progress2 --ipv4 \
	--recursive --times --links --safe-links --hard-links \
	--stats \
	--exclude "Packages*" --exclude "Sources*" \
	--exclude "Release*" --exclude "InRelease" \
	${RSYNCSOURCE} ${BASEDIR} || fatal "First stage of sync failed."

rsync --archive --human-readable --info=progress2 --ipv4 \
	--recursive --times --links --safe-links --hard-links \
	--stats --delete --delete-after \
	${RSYNCSOURCE} ${BASEDIR} || fatal "Second stage of sync failed."

date -u >${BASEDIR}/project/trace/$(hostname -f)
