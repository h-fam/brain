load("@io_bazel_rules_docker//container:push.bzl", "impl")

load("@bazel_skylib//lib:dicts.bzl", "dicts")


load(
    "@io_bazel_rules_docker//container:layer_tools.bzl",
    _get_layers = "get_from_target",
    _layer_tools = "tools",
)

load(
    "@io_bazel_rules_docker//skylib:path.bzl",
    "runfile",
)


def _get_runfile_path(ctx, f):
    return "${RUNFILES}/%s" % runfile(ctx, f)

def _push(ctx):
  print (ctx.attr)
  image = _get_layers(ctx, ctx.label.name, ctx.attr.image)
  config = image["config"]
  print (_get_runfile_path(ctx, config))
  out = ctx.actions.declare_file("%s_version.txt" % ctx.attr.name)
  print (out.path)

  cmd = "OUT=$(pwd)/%s;cd /home/marcus/src/brain;git log > ${OUT}" % (out.path)
  ctx.actions.run_shell(
    outputs = [out],
    command = cmd,
    progress_message = "Generating version in %s" % out,
  )
  filename = out.path
  output = impl(ctx)
  return output

my_push = rule (
  attrs = dicts.add({
        "format": attr.string(
            mandatory = True,
            values = [
                "OCI",
                "Docker",
            ],
            doc = "The form to push: Docker or OCI.",
        ),
        "image": attr.label(
            allow_single_file = [".tar"],
            mandatory = True,
            doc = "The label of the image to push.",
        ),
        "registry": attr.string(
            mandatory = True,
            doc = "The registry to which we are pushing.",
        ),
        "repository": attr.string(
            mandatory = True,
            doc = "The name of the image.",
        ),
        "stamp": attr.bool(
            default = False,
            mandatory = False,
        ),
        "tag": attr.string(
            default = "latest",
            doc = "(optional) The tag of the image, default to 'latest'.",
        ),
        "tag_file": attr.label(
            allow_single_file = True,
            doc = "(optional) The label of the file with tag value. Overrides 'tag'.",
        ),
        "_digester": attr.label(
            default = "@containerregistry//:digester",
            cfg = "host",
            executable = True,
        ),
        "_pusher": attr.label(
            default = Label("@containerregistry//:pusher"),
            cfg = "host",
            executable = True,
            allow_files = True,
        ),
        "_tag_tpl": attr.label(
            default = Label("@io_bazel_rules_docker//container:push-tag.sh.tpl"),
            allow_single_file = True,
        ),
    }, _layer_tools),
  implementation = _push,
  executable = True,
  toolchains = ["@io_bazel_rules_docker//toolchains/docker:toolchain_type"],
  outputs = {
    "digest": "%{name}.digest",
  },
)