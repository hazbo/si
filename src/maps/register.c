// wrengo v0.0.1-dev
//
// (c) Harry Lawrence 2015
//
// @package wrengo
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

#include <string.h>
#include "register.h"
#include "markdown.h"
#include "strings.h"

WrenForeignMethodFn findForeignMethods( WrenVM* vm, const char* module,
  const char* className, bool isStatic, const char* signature )
{
  if (is_markdown_parse(className, signature) == 1) { return class_markdown_parse; }
  if (is_strings_contains(className, signature) == 1) { return class_strings_contains; }

  if ( strcmp( className, "File" ) == 0 ) {
    if ( strcmp( signature, "Exists(_)" ) == 0 ) {
      return class_file_exists;
    }
    if ( strcmp( signature, "Read(_)" ) == 0 ) {
      return class_file_read;
    }
    if ( strcmp( signature, "Write(_,_,_)" ) == 0 ) {
      return class_file_write;
    }
  }
  return NULL;
}