#!/usr/bin/env python

import argparse
import sys

def get_args_parser():
  parser = argparse.ArgumentParser(add_help=False)
  parser.add_argument(
    "-h", "--host",
    default="localhost",
    nargs='?',
    type=str,
    help="Connect to host.")
  parser.add_argument(
    "-p", "--port",
    default=9200,
    nargs='?',
    type=int,
    help="Port number to use for connection.")
  parser.add_argument(
    "-u", "--user",
    default=None,
    nargs='?',
    type=str,
    help="Username")
  parser.add_argument(
    "-v", "--verbose",
    default=False,
    action='store_true',
    help="Print generated data"
  )
  parser.add_argument(
    "--help",
    default=False,
    action='store_true',
    help="Show this help"
  )
  return parser

if __name__ == '__main__':
  parser = get_args_parser()
  args = parser.parse_args()
  # show help by default
  if args.help or not args.user:
    parser.print_help()
    parser.exit()
    sys.exit()
  print "host " + args.host
  print "port " + str(args.port)
  print "user " + args.user
  sys.exit()
