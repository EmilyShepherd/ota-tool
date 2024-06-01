CFLAGS=\
	-I . \
	-I build \
	-I brotli/c/include \
	-I bsdiff/include \
	-I puffin/src/include \
	-I zucchini/aosp/include \
	-I libchrome \
	-I protobuf/src \
	-I protobuf/third_party/abseil-cpp

CXXFLAGS=-std=gnu++17 $(CFLAGS)
LDFLAGS=-Wl,--copy-dt-needed-entries

libs: build/combined.a build/libprotobuf-lite.a build/third_party.a

BSPATCH_SRC=\
	bsdiff/brotli_decompressor.cc \
	bsdiff/bspatch.cc \
	bsdiff/bz2_decompressor.cc \
	bsdiff/buffer_file.cc \
	bsdiff/decompressor_interface.cc \
	bsdiff/extents.cc \
	bsdiff/extents_file.cc \
	bsdiff/file.cc \
	bsdiff/logging.cc \
	bsdiff/memory_file.cc \
	bsdiff/patch_reader.cc \
	bsdiff/sink_file.cc \
	bsdiff/utils.cc

BZ2_SRC=\
	bzip2/blocksort.c \
	bzip2/huffman.c \
	bzip2/crctable.c \
	bzip2/randtable.c \
	bzip2/compress.c \
	bzip2/decompress.c \
	bzip2/bzlib.c

PUFFPATCH_SRC=\
	puffin/src/bit_reader.cc \
	puffin/src/bit_writer.cc \
	puffin/src/brotli_util.cc \
	puffin/src/huffer.cc \
	puffin/src/huffman_table.cc \
	puffin/src/memory_stream.cc \
	puffin/src/puff_reader.cc \
	puffin/src/puff_writer.cc \
	puffin/src/puffer.cc \
	puffin/src/puffin_stream.cc \
	puffin/src/puffpatch.cc \
	puffin/src/puffin.pb.cc

ZUCCHINI_SRC=\
	zucchini/abs32_utils.cc \
	zucchini/address_translator.cc \
	zucchini/arm_utils.cc \
	zucchini/binary_data_histogram.cc \
	zucchini/buffer_sink.cc \
	zucchini/buffer_source.cc \
	zucchini/crc32.cc \
	zucchini/disassembler.cc \
	zucchini/disassembler_dex.cc \
	zucchini/disassembler_elf.cc \
	zucchini/disassembler_no_op.cc \
	zucchini/disassembler_win32.cc \
	zucchini/disassembler_ztf.cc \
	zucchini/element_detection.cc \
	zucchini/encoded_view.cc \
	zucchini/ensemble_matcher.cc \
	zucchini/equivalence_map.cc \
	zucchini/heuristic_ensemble_matcher.cc \
	zucchini/image_index.cc \
	zucchini/imposed_ensemble_matcher.cc \
	zucchini/io_utils.cc \
	zucchini/patch_reader.cc \
	zucchini/patch_writer.cc \
	zucchini/reference_bytes_mixer.cc \
	zucchini/reference_set.cc \
	zucchini/rel32_finder.cc \
	zucchini/rel32_utils.cc \
	zucchini/reloc_elf.cc \
	zucchini/reloc_win32.cc \
	zucchini/target_pool.cc \
	zucchini/targets_affinity.cc \
	zucchini/zucchini_apply.cc \
	zucchini/zucchini_gen.cc \
	zucchini/zucchini_tools.cc

BROTLI_SRC=$(wildcard \
	brotli/c/common/*.c \
	brotli/c/dec/*.c \
	brotli/c/enc/*.c \
)

CHROME_SRC=\
	libchrome/base/allocator/allocator_shim_default_dispatch_to_glibc.cc \
	libchrome/base/allocator/allocator_shim_default_dispatch_to_linker_wrapped_symbols.cc \
	libchrome/base/at_exit.cc \
	libchrome/base/barrier_closure.cc \
	libchrome/base/base64_decode_fuzzer.cc \
	libchrome/base/base64_encode_fuzzer.cc \
	libchrome/base/base_paths.cc \
	libchrome/base/base_paths_posix.cc \
	libchrome/base/base_switches.cc \
	libchrome/base/big_endian.cc \
	libchrome/base/build_time.cc \
	libchrome/base/callback_helpers.cc \
	libchrome/base/callback_internal.cc \
	libchrome/base/check_example.cc \
	libchrome/base/command_line.cc \
	libchrome/base/cpu.cc \
	libchrome/base/debug/activity_analyzer.cc \
	libchrome/base/debug/activity_tracker.cc \
	libchrome/base/debug/alias.cc \
	libchrome/base/debug/asan_invalid_access.cc \
	libchrome/base/debug/crash_logging.cc \
	libchrome/base/debug/debugger.cc \
	libchrome/base/debug/debugger_posix.cc \
	libchrome/base/debug/dump_without_crashing.cc \
	libchrome/base/debug/elf_reader_linux.cc \
	libchrome/base/debug/proc_maps_linux.cc \
	libchrome/base/debug/stack_trace.cc \
	libchrome/base/debug/stack_trace_posix.cc \
	libchrome/base/debug/thread_heap_usage_tracker.cc \
	libchrome/base/deferred_sequenced_task_runner.cc \
	libchrome/base/environment.cc \
	libchrome/base/feature_list.cc \
	libchrome/base/file_descriptor_store.cc \
	libchrome/base/files/file.cc \
	libchrome/base/files/file_descriptor_watcher_posix.cc \
	libchrome/base/files/file_enumerator.cc \
	libchrome/base/files/file_enumerator_posix.cc \
	libchrome/base/files/file_path.cc \
	libchrome/base/files/file_path_constants.cc \
	libchrome/base/files/file_path_watcher.cc \
	libchrome/base/files/file_path_watcher_linux.cc \
	libchrome/base/files/file_path_watcher_stub.cc \
	libchrome/base/files/file_posix.cc \
	libchrome/base/files/file_proxy.cc \
	libchrome/base/files/file_tracing.cc \
	libchrome/base/files/file_util.cc \
	libchrome/base/files/file_util_linux.cc \
	libchrome/base/files/file_util_posix.cc \
	libchrome/base/files/important_file_writer.cc \
	libchrome/base/files/memory_mapped_file.cc \
	libchrome/base/files/memory_mapped_file_posix.cc \
	libchrome/base/files/scoped_file.cc \
	libchrome/base/files/scoped_temp_dir.cc \
	libchrome/base/guid.cc \
	libchrome/base/hash.cc \
	libchrome/base/json/json_correctness_fuzzer.cc \
	libchrome/base/json/json_file_value_serializer.cc \
	libchrome/base/json/json_parser.cc \
	libchrome/base/json/json_reader.cc \
	libchrome/base/json/json_reader_fuzzer.cc \
	libchrome/base/json/json_string_value_serializer.cc \
	libchrome/base/json/json_value_converter.cc \
	libchrome/base/json/json_writer.cc \
	libchrome/base/json/string_escape.cc \
	libchrome/base/json/string_escape_fuzzer.cc \
	libchrome/base/lazy_instance_helpers.cc \
	libchrome/base/linux_util.cc \
	libchrome/base/location.cc \
	libchrome/base/logging.cc \
	libchrome/base/md5.cc \
	libchrome/base/memory/aligned_memory.cc \
	libchrome/base/memory/discardable_memory_allocator.cc \
	libchrome/base/memory/discardable_memory.cc \
	libchrome/base/memory/discardable_shared_memory.cc \
	libchrome/base/memory/memory_coordinator_client.cc \
	libchrome/base/memory/memory_coordinator_client_registry.cc \
	libchrome/base/memory/memory_coordinator_proxy.cc \
	libchrome/base/memory/memory_pressure_listener.cc \
	libchrome/base/memory/memory_pressure_monitor.cc \
	libchrome/base/memory/memory_pressure_monitor_chromeos.cc \
	libchrome/base/memory/platform_shared_memory_region.cc \
	libchrome/base/memory/platform_shared_memory_region_posix.cc \
	libchrome/base/memory/protected_memory.cc \
	libchrome/base/memory/protected_memory_posix.cc \
	libchrome/base/memory/read_only_shared_memory_region.cc \
	libchrome/base/memory/ref_counted.cc \
	libchrome/base/memory/ref_counted_memory.cc \
	libchrome/base/memory/shared_memory_handle.cc \
	libchrome/base/memory/shared_memory_handle_posix.cc \
	libchrome/base/memory/shared_memory_helper.cc \
	libchrome/base/memory/shared_memory_mapping.cc \
	libchrome/base/memory/shared_memory_posix.cc \
	libchrome/base/memory/unsafe_shared_memory_region.cc \
	libchrome/base/memory/weak_ptr.cc \
	libchrome/base/memory/writable_shared_memory_region.cc \
	libchrome/base/message_loop/incoming_task_queue.cc \
	libchrome/base/message_loop/message_loop.cc \
	libchrome/base/message_loop/message_loop_current.cc \
	libchrome/base/message_loop/message_loop_task_runner.cc \
	libchrome/base/message_loop/message_pump.cc \
	libchrome/base/message_loop/message_pump_default.cc \
	libchrome/base/message_loop/message_pump_libevent.cc \
	libchrome/base/message_loop/watchable_io_message_pump_posix.cc \
	libchrome/base/metrics/bucket_ranges.cc \
	libchrome/base/metrics/dummy_histogram.cc \
	libchrome/base/metrics/field_trial.cc \
	libchrome/base/metrics/field_trial_param_associator.cc \
	libchrome/base/metrics/field_trial_params.cc \
	libchrome/base/metrics/histogram_base.cc \
	libchrome/base/metrics/histogram.cc \
	libchrome/base/metrics/histogram_delta_serialization.cc \
	libchrome/base/metrics/histogram_functions.cc \
	libchrome/base/metrics/histogram_samples.cc \
	libchrome/base/metrics/histogram_snapshot_manager.cc \
	libchrome/base/metrics/metrics_hashes.cc \
	libchrome/base/metrics/persistent_histogram_allocator.cc \
	libchrome/base/metrics/persistent_histogram_storage.cc \
	libchrome/base/metrics/persistent_memory_allocator.cc \
	libchrome/base/metrics/persistent_sample_map.cc \
	libchrome/base/metrics/sample_map.cc \
	libchrome/base/metrics/sample_vector.cc \
	libchrome/base/metrics/single_sample_metrics.cc \
	libchrome/base/metrics/sparse_histogram.cc \
	libchrome/base/metrics/statistics_recorder.cc \
	libchrome/base/metrics/user_metrics.cc \
	libchrome/base/native_library.cc \
	libchrome/base/native_library_posix.cc \
	libchrome/base/observer_list_threadsafe.cc \
	libchrome/base/path_service.cc \
	libchrome/base/pending_task.cc \
	libchrome/base/pickle.cc \
	libchrome/base/posix/file_descriptor_shuffle.cc \
	libchrome/base/posix/global_descriptors.cc \
	libchrome/base/posix/safe_strerror.cc \
	libchrome/base/posix/unix_domain_socket.cc \
	libchrome/base/power_monitor/power_monitor.cc \
	libchrome/base/power_monitor/power_monitor_device_source.cc \
	libchrome/base/power_monitor/power_monitor_device_source_chromeos.cc \
	libchrome/base/power_monitor/power_monitor_device_source_stub.cc \
	libchrome/base/power_monitor/power_monitor_source.cc \
	libchrome/base/process/internal_linux.cc \
	libchrome/base/process/kill.cc \
	libchrome/base/process/kill_posix.cc \
	libchrome/base/process/launch.cc \
	libchrome/base/process/launch_posix.cc \
	libchrome/base/process/memory_stubs.cc \
	libchrome/base/process/process_handle.cc \
	libchrome/base/process/process_handle_linux.cc \
	libchrome/base/process/process_handle_posix.cc \
	libchrome/base/process/process_info_linux.cc \
	libchrome/base/process/process_iterator.cc \
	libchrome/base/process/process_iterator_linux.cc \
	libchrome/base/process/process_linux.cc \
	libchrome/base/process/process_metrics.cc \
	libchrome/base/process/process_metrics_linux.cc \
	libchrome/base/process/process_metrics_posix.cc \
	libchrome/base/process/process_posix.cc \
	libchrome/base/profiler/native_stack_sampler.cc \
	libchrome/base/profiler/native_stack_sampler_posix.cc \
	libchrome/base/profiler/stack_sampling_profiler.cc \
	libchrome/base/profiler/test_support_library.cc \
	libchrome/base/rand_util.cc \
	libchrome/base/rand_util_posix.cc \
	libchrome/base/run_loop.cc \
	libchrome/base/scoped_native_library.cc \
	libchrome/base/sequence_checker_impl.cc \
	libchrome/base/sequenced_task_runner.cc \
	libchrome/base/sequence_token.cc \
	libchrome/base/sha1.cc \
	libchrome/base/strings/latin1_string_conversions.cc \
	libchrome/base/strings/nullable_string16.cc \
	libchrome/base/strings/old_utf_string_conversions.cc \
	libchrome/base/strings/pattern.cc \
	libchrome/base/strings/safe_sprintf.cc \
	libchrome/base/strings/strcat.cc \
	libchrome/base/strings/string16.cc \
	libchrome/base/strings/string_number_conversions.cc \
	libchrome/base/strings/string_number_conversions_fuzzer.cc \
	libchrome/base/strings/string_piece.cc \
	libchrome/base/strings/stringprintf.cc \
	libchrome/base/strings/string_split.cc \
	libchrome/base/strings/string_tokenizer_fuzzer.cc \
	libchrome/base/strings/string_util.cc \
	libchrome/base/strings/string_util_constants.cc \
	libchrome/base/strings/sys_string_conversions_posix.cc \
	libchrome/base/strings/utf_offset_string_conversions.cc \
	libchrome/base/strings/utf_string_conversions.cc \
	libchrome/base/strings/utf_string_conversions_fuzzer.cc \
	libchrome/base/strings/utf_string_conversions_regression_fuzzer.cc \
	libchrome/base/strings/utf_string_conversion_utils.cc \
	libchrome/base/supports_user_data.cc \
	libchrome/base/synchronization/atomic_flag.cc \
	libchrome/base/synchronization/condition_variable_posix.cc \
	libchrome/base/synchronization/lock.cc \
	libchrome/base/synchronization/lock_impl_posix.cc \
	libchrome/base/synchronization/waitable_event_posix.cc \
	libchrome/base/synchronization/waitable_event_watcher_posix.cc \
	libchrome/base/sync_socket_posix.cc \
	libchrome/base/sys_info.cc \
	libchrome/base/sys_info_chromeos.cc \
	libchrome/base/sys_info_linux.cc \
	libchrome/base/sys_info_posix.cc \
	libchrome/base/syslog_logging.cc \
	libchrome/base/system_monitor/system_monitor.cc \
	libchrome/base/task/cancelable_task_tracker.cc \
	libchrome/base/task_runner.cc \
	libchrome/base/third_party/icu/icu_utf.cc \
	libchrome/base/third_party/nspr/prtime.cc \
	libchrome/base/threading/platform_thread_internal_posix.cc \
	libchrome/base/threading/platform_thread_linux.cc \
	libchrome/base/threading/platform_thread_posix.cc \
	libchrome/base/threading/post_task_and_reply_impl.cc \
	libchrome/base/threading/scoped_blocking_call.cc \
	libchrome/base/threading/sequenced_task_runner_handle.cc \
	libchrome/base/threading/sequence_local_storage_map.cc \
	libchrome/base/threading/sequence_local_storage_slot.cc \
	libchrome/base/threading/simple_thread.cc \
	libchrome/base/threading/thread.cc \
	libchrome/base/threading/thread_checker_impl.cc \
	libchrome/base/threading/thread_collision_warner.cc \
	libchrome/base/threading/thread_id_name_manager.cc \
	libchrome/base/threading/thread_local_storage.cc \
	libchrome/base/threading/thread_local_storage_posix.cc \
	libchrome/base/threading/thread_restrictions.cc \
	libchrome/base/threading/thread_task_runner_handle.cc \
	libchrome/base/threading/watchdog.cc \
	libchrome/base/time/clock.cc \
	libchrome/base/time/default_clock.cc \
	libchrome/base/time/default_tick_clock.cc \
	libchrome/base/timer/elapsed_timer.cc \
	libchrome/base/timer/hi_res_timer_manager_posix.cc \
	libchrome/base/timer/mock_timer.cc \
	libchrome/base/timer/timer.cc \
	libchrome/base/time/tick_clock.cc \
	libchrome/base/time/time.cc \
	libchrome/base/time/time_conversion_posix.cc \
	libchrome/base/time/time_exploded_posix.cc \
	libchrome/base/time/time_now_posix.cc \
	libchrome/base/time/time_override.cc \
	libchrome/base/time/time_to_iso8601.cc \
	libchrome/base/unguessable_token.cc \
	libchrome/base/value_conversions.cc \
	libchrome/base/value_iterators.cc \
	libchrome/base/values.cc \
	libchrome/base/version.cc \
	libchrome/base/vlog.cc \

ALL_SRC=\
	$(BSPATCH_SRC) \
	$(BZ2_SRC) \
	$(PUFFPATCH_SRC) \
	$(ZUCCHINI_SRC) \
	$(BROTLI_SRC) \
	$(CHROME_SRC)

ALL_OBJ=$(addprefix build/,$(patsubst %.cc,%.o,$(patsubst %.c,%.o, $(ALL_SRC))))

# Need to replace the LOG statement with a single ; so that a statement
# still exists. This is because some LOG calls are in an if/else block
# without brackets.
build/%.cc: %.cc
	@ mkdir -p $(dir $@)
	@ sed -z 's/\(\n *\)[A-Z_]*LOG[^;]*;\n/\1;\n/g' $< > $@

build/%.o: build/%.cc
	@ $(CXX) $(CXXFLAGS) -c -o $@ $< $(LDFLAGS)
	@ echo [CPP] $@

build/%.o: %.c
	@ mkdir -p $(dir $@)
	@ $(CC) $(CFLAGS) -c -o $@ $< $(LDFLAGS)
	@ echo [CC] $@

build/combined.a: $(ALL_OBJ)
	@ $(AR) rcs $@ $^
	@ echo [AR] $@

puffin/src/puffpatch.cc: build/puffin/src/puffin.pb.cc
build/puffin/src/puffin.pb.cc: puffin/src/puffin.proto
	protoc -I=puffin/src/ --cpp_out=build/puffin/src/ $^

build/third_party.a: build/protobuf/libprotobuf-lite.a
	@ find build/protobuf/third_party -name '*.a' | xargs $(AR) -rcT $@
	@ echo [AR] $@

build/libprotobuf-lite.a: build/protobuf/libprotobuf-lite.a
	cp $< $@

build/protobuf/libprotobuf-lite.a: build/protobuf/Makefile
	make -C build/protobuf libprotobuf-lite

build/protobuf/Makefile:
	cmake -S protobuf -B build/protobuf

ALL_OBJ=build/
clean:
	rm -rf $(ALL_OBJ)
